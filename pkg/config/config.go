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

type AWS struct {
	SignatureBucket       string `env:"SIGNATURE_BUCKET,required"`
	SignatureBucketRegion string `env:"SIGNATURE_BUCKET_REGION,required"`
	AccessKey             string `env:"AWS_ACCESS_KEY,required"`
	SecretKey             string `env:"AWS_SECRET_ACCESS_KEY, required"`
}

func (d DB) GetDSN() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		d.User,
		d.Password,
		d.Name,
		d.Port)
}

type App struct {
	Server Server
	DB
	AWS
}

var Config App

func init() {
	ctx := context.TODO()

	if err := envconfig.Process(ctx, &Config); err != nil {
		panic(err)
	}
}
