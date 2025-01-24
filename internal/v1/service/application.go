package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/cromerc/projectBase/internal/v1/domain"

	"github.com/cromerc/projectBase/internal/v1/port/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTTPServer struct {
	ctx         context.Context
	log         domain.Logger
	server      api.StrictServerInterface
	startupTime time.Time
}

func NewHTTPServer(ctx context.Context, logger domain.Logger, server api.StrictServerInterface) *HTTPServer {
	logger.Ctx(ctx)
	return &HTTPServer{
		ctx:         ctx,
		log:         logger,
		server:      server,
		startupTime: time.Now().UTC(),
	}
}

func (hs HTTPServer) Run() (err error) {
	r := chi.NewRouter()

	apiRouter := chi.NewRouter()

	strictMode := api.NewStrictHandlerWithOptions(hs.server, nil, configOptions())
	api.HandlerFromMux(strictMode, apiRouter)

	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json", "application/xml"))
	r.Use(middleware.Compress(5))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Throttle(3))
	r.Use(middleware.Timeout(time.Minute))

	r.Mount("/v1", apiRouter)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		hs.log.Info("Starting HTTP server")
		if err = httpServer.ListenAndServe(); !errors.Is(http.ErrServerClosed, err) {
			hs.log.Fatal("failed to start http server", err)
		}
	}()
	<-hs.ctx.Done()
	hs.log.Info("Stopping HTTP server")
	if err = httpServer.Shutdown(context.TODO()); err != nil {
		hs.log.Fatal("failed to shutdown http server", err)
	}

	return
}

func configOptions() api.StrictHTTPServerOptions {
	return api.StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
		},
	}
}
