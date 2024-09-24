package main

import (
	"log"

	"github.com/incwadi-warehouse/monorepo-go/blog/router"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
