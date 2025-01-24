package mapper

import (
	"github.com/cromerc/projectBase/internal/v1/domain"
	"github.com/cromerc/projectBase/internal/v1/port/api"
)

//go:generate go run -modfile=../../../../tools/go.mod github.com/jmattheis/goverter/cmd/goverter gen ./

// goverter:converter
// goverter:output:file ./mapper.gen.go
// goverter:output:package github.com/cromerc/projectBase/internal/v1/port/mapper
// goverter:name PortMapper
// goverter:extend strconv:Atoi
// goverter:extend strconv:Itoa
// goverter:extend userPortToDomain
// goverter:matchIgnoreCase
// goverter:wrapErrors
type Mapper interface {
	// goverter:map Username Name
	UserDomainToPort(source domain.User) api.User
	UsersDomainToPort(source []domain.User) []api.User
	UserPortToDomain(source api.User) (domain.User, error)
	UsersPortToDomain(source []api.User) ([]domain.User, error)
}

func userPortToDomain(_ Mapper, source api.User) (domain.User, error) {
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
