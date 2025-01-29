package api

import (
	"github.com/cromerc/projectBase/internal/v1/domain"
	"github.com/cromerc/projectBase/internal/v1/port/api"
)

//go:generate go run -modfile=../../../../../tools/go.mod github.com/jmattheis/goverter/cmd/goverter gen ./

// goverter:converter
// goverter:output:file ./domain_to_api.gen.go
// goverter:output:package github.com/cromerc/projectBase/internal/v1/port/mapper/api
// goverter:name toAPI
// goverter:extend strconv:Atoi
// goverter:extend strconv:Itoa
// goverter:matchIgnoreCase
// goverter:wrapErrors
type ToAPI interface {
	// goverter:map Username Name
	User(source domain.User) api.User
	Users(source []domain.User) []api.User
}

// goverter:converter
// goverter:output:file ./api_to_domain.gen.go
// goverter:output:package github.com/cromerc/projectBase/internal/v1/port/mapper/api
// goverter:name toDomain
// goverter:extend strconv:Atoi
// goverter:extend strconv:Itoa
// goverter:extend github.com/cromerc/projectBase/internal/v1/port/mapper/api:NewDomainUser
// goverter:matchIgnoreCase
// goverter:wrapErrors
type ToDomain interface {
	User(source api.User) (domain.User, error)
	Users(source []api.User) ([]domain.User, error)
}

func NewDomainUser(_ ToDomain, source api.User) (domain.User, error) {
	id := int64(0)

	if source.ID != nil {
		id = *source.ID
	}

	user, err := domain.NewUser(id, source.Name)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
