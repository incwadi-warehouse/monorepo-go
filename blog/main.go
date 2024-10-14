package main

import (
	"log"

	"github.com/incwadi-warehouse/monorepo-go/blog/router"
	"github.com/incwadi-warehouse/monorepo-go/framework/config"
	"github.com/incwadi-warehouse/monorepo-go/framework/cors"
)

func main() {
	config.LoadAppConfig()

	corsConfig := cors.NewCors()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.SetCorsHeaders()

	r := router.Routes()
	r.Use(corsConfig.SetCorsHeaders())

	log.Fatal(r.Run(":8080"))
}
