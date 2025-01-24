package port

import (
	"context"

	"github.com/cromerc/projectBase/internal/v1/port/mapper"

	"github.com/cromerc/projectBase/internal/v1/port/api"
)

//go:generate go run -modfile=../../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config/server.yaml ../../../api/v1.yaml
//go:generate go run -modfile=../../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config/models.yaml ../../../api/v1.yaml
//go:generate go run -modfile=../../../tools/go.mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=config/client.yaml ../../../api/v1.yaml

type server struct {
	mapper mapper.PortMapper
}

func (s server) GetAllTODOTasks(ctx context.Context, request api.GetAllTODOTasksRequestObject) (api.GetAllTODOTasksResponseObject, error) {
	// TODO implement me
	return api.GetAllTODOTasks501Response{}, nil
}

func (s server) CreateTODO(ctx context.Context, request api.CreateTODORequestObject) (api.CreateTODOResponseObject, error) {
	// TODO implement me
	return api.CreateTODO501Response{}, nil
}

func (s server) DeleteUser(ctx context.Context, request api.DeleteUserRequestObject) (api.DeleteUserResponseObject, error) {
	// TODO implement me
	return api.DeleteUser501Response{}, nil
}

func (s server) GetUserByName(ctx context.Context, request api.GetUserByNameRequestObject) (api.GetUserByNameResponseObject, error) {
	// TODO implement me
	return api.GetUserByName501Response{}, nil
}

func (s server) UpdateUser(ctx context.Context, request api.UpdateUserRequestObject) (api.UpdateUserResponseObject, error) {
	// TODO implement me
	return api.UpdateUser501Response{}, nil
}

func NewServer() api.StrictServerInterface {
	pm := mapper.PortMapper{}

	return server{mapper: pm}
}
