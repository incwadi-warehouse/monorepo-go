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
	case "add":
		s := settings.Load(os.Args[1], os.Args[2])
		s.Add(os.Args[4], os.Args[5])
		s.Write()
	case "rem":
		s := settings.Load(os.Args[1], os.Args[2])
		s.Rem(os.Args[4])
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
	fmt.Println("add - Add new entry")
	fmt.Println("rem - Remove an entry")
}
