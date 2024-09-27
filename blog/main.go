package main

import (
	"log"

	"github.com/incwadi-warehouse/monorepo-go/blog/router"
	"github.com/incwadi-warehouse/monorepo-go/framework/config"
)

func main() {
    config.LoadAppConfig()

	r := router.Routes()
	log.Fatal(r.Run(":8080"))
}
