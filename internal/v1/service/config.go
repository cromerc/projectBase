package service

//go:generate go run -modfile=../../../tools/go.mod github.com/g4s8/envdoc -format markdown -required-if-no-def -output ../../../env.md -types [!Config]*

type Config struct {
	Log      Log      `envPrefix:"LOG_"`
	Database Database `envPrefix:"DB_"`
}

type Log struct {
	Level  string `env:"LEVEL" envDefault:"debug"`
	Pretty bool   `env:"PRETTY" envDefault:"true"`
}

type Database struct {
	ReadHost  string `env:"READ_HOST,required,notEmpty"`
	WriteHost string `env:"WRITE_HOST,required,notEmpty"`
	Port      string `env:"PORT,required,notEmpty"`
	Username  string `env:"USERNAME,required,notEmpty"`
	Password  string `env:"PASSWORD,required,notEmpty"`
	Database  string `env:"DATABASE,required,notEmpty"`
	Schema    string `env:"SCHEMA,required,notEmpty"`
}
