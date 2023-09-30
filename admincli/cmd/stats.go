package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/bcicen/go-units"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show stats",
	Long:  `Show stats`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := os.ReadDir(viper.GetString("stats_dir"))
		if err != nil {
			log.Fatalln(err)
			return
		}

		var totalSize int64
		for _, file := range files {
			if !file.IsDir() {
				fileinfo, _ := file.Info()
				totalSize += fileinfo.Size()
			}
		}

		size := units.NewValue(float64(totalSize), units.Byte)
		convertedSize := size.MustConvert(units.Gibibyte)
		opts := units.FmtOptions{
			Label:     true,
			Short:     true,
			Precision: 2,
		}

		fmt.Println("Directory size")
		fmt.Println("Path", viper.GetString("stats_dir"))
		fmt.Println("Size", convertedSize.Fmt(opts))
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
