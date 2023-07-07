package cmd

import (
	"fmt"
	"log"

	"github.com/incwadi-warehouse/monorepo-go/admincli/command"
	"github.com/spf13/cobra"
)

var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot the OS.",
	Long:  `Reboot the OS. The connection will be terminated.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Reboot system")
		fmt.Println("The connection will be terminated")

		out, err := command.Command([]string{"reboot"})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(out)
	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)
}
