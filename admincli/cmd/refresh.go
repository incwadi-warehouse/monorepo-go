package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/incwadi-warehouse/monorepo-go/admincli/command"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Reloads container images",
	Long:  `Restart all containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping all containers...")

		stopContainers, err := command.Command([]string{"/usr/bin/docker", "compose", "--project-directory", viper.GetString("project_dir"), "down"})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(stopContainers))

		fmt.Println("Starting all containers...")

		startContainers, err := command.Command([]string{"docker compose --project-directory " + os.Getenv("PROJECT_DIR") + " up -d"})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(startContainers))
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}
