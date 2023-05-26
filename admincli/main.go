package main

import (
	"log"
	"os"

	"github.com/incwadi-warehouse/monorepo-go/admincli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat("./.env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cmd.Execute()
}
