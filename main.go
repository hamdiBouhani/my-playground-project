package main

import (
	"log"

	"github.com/hamdiBouhani/my-playground-project/cmd"
	"github.com/joho/godotenv"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
