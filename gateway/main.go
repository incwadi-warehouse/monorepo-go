package main

import (
	"github.com/incwadi-warehouse/monorepo-go/gateway/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router.Router()
}
