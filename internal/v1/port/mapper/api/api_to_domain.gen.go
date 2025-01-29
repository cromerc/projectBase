// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package api

import (
	"fmt"
	domain "github.com/cromerc/projectBase/internal/v1/domain"
	api "github.com/cromerc/projectBase/internal/v1/port/api"
)

type toDomain struct{}

func (c *toDomain) User(source api.User) (domain.User, error) {
	return NewDomainUser(c, source)
}
func (c *toDomain) Users(source []api.User) ([]domain.User, error) {
	var domainUserList []domain.User
	if source != nil {
		domainUserList = make([]domain.User, len(source))
		for i := 0; i < len(source); i++ {
			domainUser, err := NewDomainUser(c, source[i])
			if err != nil {
				return nil, fmt.Errorf("error setting index %d: %w", i, err)
			}
			domainUserList[i] = domainUser
		}
	}
	return domainUserList, nil
}
