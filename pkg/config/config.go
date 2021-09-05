package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Server struct {
	Port int `env:"SERVER_PORT,required"`
}

type DB struct {
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Port     string `env:"DB_PORT,required"`
	Name     string `env:"DB_NAME,required "`
}

func (d DB) GetDSN() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", d.User, d.Password, d.Name, d.Port)
}

type App struct {
	Server Server
	DB
}

var Config App

func init() {
	ctx := context.TODO()

	if err := envconfig.Process(ctx, &Config); err != nil {
		panic(err)
	}
}
