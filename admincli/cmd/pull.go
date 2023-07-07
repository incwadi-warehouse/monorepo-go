package cmd

import (
	"fmt"
	"log"

	"github.com/incwadi-warehouse/monorepo-go/admincli/command"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Refresh container images",
	Long:  `Fetch the latest images from the registry. After restarting the container the new image will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pulling new images...")

		out, err := command.Command([]string{"/usr/bin/docker compose --project-directory "+ viper.GetString("project_dir") +" pull"})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(out)
	},
}


func init() {
	rootCmd.AddCommand(pullCmd)
}
