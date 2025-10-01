package main

import (
	"github.com/joho/godotenv"

	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	app := application{
		baseURL: os.Getenv("BASE_URL"),
		token:   os.Getenv("TOKEN"),
	}

	err = app.reqAuthentication()
	if err != nil {
		logger.Error(err.Error())
		return
	}
}
