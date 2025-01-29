package mapper

import "github.com/cromerc/projectBase/internal/v1/adapter/mapper/postgres"

type Postgres struct {
	ToDomain   postgres.ToDomain
	ToPostgres postgres.ToPostgres
}

func NewPostgres() Postgres {
	return Postgres{
		ToDomain:   postgres.NewToDomain(),
		ToPostgres: postgres.NewToPostgres(),
	}
}
