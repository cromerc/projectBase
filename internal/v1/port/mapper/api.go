package mapper

import (
	"github.com/cromerc/projectBase/internal/v1/port/mapper/api"
)

type API struct {
	ToDomain api.ToDomain
	ToAPI    api.ToAPI
}

func NewAPI() API {
	return API{
		ToDomain: api.NewToDomain(),
		ToAPI:    api.NewToAPI(),
	}
}
