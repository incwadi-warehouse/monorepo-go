package cmd

import (
	"fmt"
	"log"

	"github.com/incwadi-warehouse/monorepo-go/admincli/command"
	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Reloads container images",
	Long:  `Restart all containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping all containers...")

		stopContainers()

		fmt.Println("Starting all containers...")

		startContainers()
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}

func stopContainers() {
	out, err := command.Command([]string{"docker compose --project-directory . down"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(out)
}

func startContainers() {
	out, err := command.Command([]string{"docker compose --project-directory . up -d"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(out)
}
