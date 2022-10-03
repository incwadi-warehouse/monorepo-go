package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/incwadi-warehouse/monorepo-go/conf/settings"
)

func main() {
	var action string
    if len(os.Args) >= 2 {
		action = os.Args[1]
	}

	switch action {
	case "get":
        if err := hasArgs(5); err != nil {
		    log.Fatal(err)
            os.Exit(1)
	    }

		s, err := settings.LoadFromUrl(os.Args[3], os.Args[4])
		if err != nil {
			log.Fatal(err)
		}

		v := s.Get(os.Args[2])
		fmt.Println(v)
	case "add":
        if err := hasArgs(6); err != nil {
		    log.Fatal(err)
            os.Exit(1)
	    }

		s, err := settings.LoadFromUrl(os.Args[4], os.Args[5])
		if err != nil {
			log.Fatal(err)
		}

		s.Add(os.Args[2], os.Args[3])
		s.Write()
	case "rm":
        if err := hasArgs(5); err != nil {
		    log.Fatal(err)
            os.Exit(1)
	    }

		s, err := settings.LoadFromUrl(os.Args[3], os.Args[4])
		if err != nil {
			log.Fatal(err)
		}

		s.Rm(os.Args[2])
		s.Write()
	case "help":
		help()
    default:
        help()
        fmt.Println("")

        log.Fatal("NO ACTION PASSED")
        os.Exit(1)
	}
}

func hasArgs(counter int) error {
    if len(os.Args) != counter {
        return errors.New("ARGUMENTS MISSING")
    }

    return nil
}

func help() {
	fmt.Println("Settings")
	fmt.Println("")
	fmt.Println("Usage: conf [action]")
	fmt.Println("")
	fmt.Println("Actions")
	fmt.Println("get [key] [value] [schema-url] [file-url] - Get the value of an entry")
	fmt.Println("add [key] [schema-url] [file-url] - Add or update entry")
	fmt.Println("rm [key] [schema-url] [file-url] - Remove an entry")
}
