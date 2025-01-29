package postgres

//go:generate go run -modfile=../../../../../tools/go.mod github.com/jmattheis/goverter/cmd/goverter gen ./

// goverter:converter
// goverter:output:file ./domain_to_postgres.gen.go
// goverter:output:package github.com/cromerc/projectBase/internal/v1/adapter/mapper/postgres
// goverter:name toPostgres
// goverter:extend strconv:Atoi
// goverter:extend strconv:Itoa
// goverter:matchIgnoreCase
// goverter:wrapErrors
type ToPostgres interface{}

// goverter:converter
// goverter:output:file ./postgres_to_domain.gen.go
// goverter:output:package github.com/cromerc/projectBase/internal/v1/adapter/mapper/postgres
// goverter:name toDomain
// goverter:extend strconv:Atoi
// goverter:extend strconv:Itoa
// goverter:matchIgnoreCase
// goverter:wrapErrors
type ToDomain interface{}
