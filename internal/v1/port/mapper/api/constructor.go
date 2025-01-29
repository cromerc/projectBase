//go:build !goverter

package api

func NewToDomain() ToDomain {
	return &toDomain{}
}

func NewToAPI() ToAPI {
	return &toAPI{}
}
