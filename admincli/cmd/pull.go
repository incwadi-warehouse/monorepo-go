package cmd

import (
	"fmt"
	"log"

	"github.com/incwadi-warehouse/monorepo-go/admincli/command"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Refresh container images",
	Long:  `Fetch the latest images from the registry. After restarting the container the new image will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pulling new images...")

		out, err := command.Command([]string{"docker compose --project-directory . pull"})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(out)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
