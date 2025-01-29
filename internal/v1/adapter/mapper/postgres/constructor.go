//go:build !goverter

package postgres

func NewToDomain() ToDomain {
	return &toDomain{}
}

func NewToPostgres() ToPostgres {
	return &toPostgres{}
}
