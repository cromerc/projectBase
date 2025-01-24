package mapper

import (
	"github.com/cromerc/projectBase/internal/v3/domain"
	"github.com/cromerc/projectBase/internal/v3/port/api"
)

//go:generate go run -modfile=../../../../tools/go.mod github.com/jmattheis/goverter/cmd/goverter gen ./

// goverter:converter
// goverter:output:file ./mapper.gen.go
// goverter:output:package github.com/cromerc/projectBase/internal/v3/port/mapper
// goverter:name PortMapper
// goverter:extend strconv:Atoi
// goverter:extend strconv:Itoa
// goverter:extend userPortToDomain
// goverter:matchIgnoreCase
// goverter:wrapErrors
type Mapper interface {
	// goverter:ignore UserStatus
	UserDomainToPort(source domain.User) api.User
	UsersDomainToPort(source []domain.User) []api.User
	UserPortToDomain(source api.User) (domain.User, error)
	UsersPortToDomain(source []api.User) ([]domain.User, error)
}

func userPortToDomain(_ Mapper, source api.User) (domain.User, error) {
	id := int64(0)
	firstName := ""
	lastName := ""
	email := ""
	password := ""

	if source.ID != nil {
		id = *source.ID
	}
	if source.FirstName != nil {
		firstName = *source.FirstName
	}
	if source.LastName != nil {
		lastName = *source.LastName
	}
	if source.Email != nil {
		email = *source.Email
	}
	if source.Password != nil {
		password = *source.Password
	}

	user, err := domain.NewUser(id, firstName, lastName, email, password)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
