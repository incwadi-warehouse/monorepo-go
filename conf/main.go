package main

import (
	"fmt"
	"log"
	"os"

	"github.com/incwadi-warehouse/monorepo-go/conf/settings"
)

func main() {
	var action interface{}
	if len(os.Args) >= 4 {
		action = os.Args[3]
	}
	if len(os.Args) == 2 && os.Args[1] == "help" {
		action = "help"
	}

	switch action {
	case "get":
		s, err := settings.LoadFromUrl(os.Args[1], os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		v := s.Get(os.Args[4])
		fmt.Println(v)
	case "add":
		s, err := settings.LoadFromUrl(os.Args[1], os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		s.Add(os.Args[4], os.Args[5])
		s.Write()
	case "rm":
		s, err := settings.LoadFromUrl(os.Args[1], os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		s.Rm(os.Args[4])
		s.Write()
	case "help":
		help()
	default:
		log.Println("Arguments missing")
		fmt.Println("")
		help()
		os.Exit(1)
	}
}

func help() {
	fmt.Println("Settings")
	fmt.Println("")
	fmt.Println("Usage: conf [schema-url] [file-url] [action]")
	fmt.Println("")
	fmt.Println("Actions")
	fmt.Println("get - Get the value of an entry")
	fmt.Println("add - Add new entry")
	fmt.Println("rm - Remove an entry")
}
