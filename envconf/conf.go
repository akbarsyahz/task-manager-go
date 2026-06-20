package envconf

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port     string
	Host     string
	Name     string
	User     string
	Pass     string
	SSL      string
	Timezone string
}

func EnvSetting() (*Env, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	env := &Env{
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASSWORD"),
		SSL:      os.Getenv("DB_SSL"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}

	return env, nil

}
