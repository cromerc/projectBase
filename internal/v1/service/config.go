package service

type Config struct {
	Log      Log      `envPrefix:"LOG_"`
	Database Database `envPrefix:"DB_"`
}

type Log struct {
	Level  string `env:"LEVEL" envDefault:"debug"`
	Pretty bool   `env:"PRETTY" envDefault:"true"`
}

type Database struct {
	ReadHost  string `env:"READ_HOST" envDefault:"localhost"`
	WriteHost string `env:"WRITE_HOST" envDefault:"localhost"`
	Port      string `env:"PORT" envDefault:"5432"`
	Username  string `env:"USERNAME" envDefault:"postgres"`
	Password  string `env:"PASSWORD" envDefault:"postgres"`
	Database  string `env:"DATABASE" envDefault:"postgres"`
	Schema    string `env:"SCHEMA" envDefault:"public"`
}
