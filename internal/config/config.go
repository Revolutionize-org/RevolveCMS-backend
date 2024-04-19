package config

import "github.com/ilyakaznacheev/cleanenv"

var Config = config{}

type config struct {
	Postgres Postgres
	Secret   secret
	Api      api
}

type Postgres struct {
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	DB       string `env:"POSTGRES_DB"`
	Port     string `env:"POSTGRES_PORT"`
	URL      string `env:"POSTGRES_URL"`
}

type secret struct {
	AccessToken  string `env:"ACCESS_TOKEN_SECRET"`
	RefreshToken string `env:"REFRESH_TOKEN_SECRET"`
}

type api struct {
	Port string `env:"API_PORT"`
	Env  string `env:"ENV"`
}

func init() {
	err := cleanenv.ReadConfig(".env", &Config)
	if err != nil {
		panic(err)
	}
}
