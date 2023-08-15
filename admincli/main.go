package main

import (
	"github.com/incwadi-warehouse/monorepo-go/admincli/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load() // .env

	cmd.Execute()
}
