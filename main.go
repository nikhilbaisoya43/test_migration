package main

import (
	"04_pro/api"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("env")
	if env == "" {
		env = "default"
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occured while loading .env file", err)
	}
	appName := os.Getenv("APP_NAME")
	appPort := os.Getenv("APP_PORT")
	fmt.Printf(fmt.Sprintf("%s starting on http://localhost:%s\n", appName, appPort))

	api.INitializeServer()
}
