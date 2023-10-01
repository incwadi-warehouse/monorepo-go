package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bcicen/go-units"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show stats",
	Long:  `Show stats`,
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Directory size")

		totalSize, err := getDirSize(viper.GetString("stats_dir"))
        if err != nil {
            log.Fatalln(err)
        }

		fmt.Println("Path", viper.GetString("stats_dir"))
		fmt.Println("Size", convertSizeToGb(totalSize))
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

func getDirSize(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(_ string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			size += fileInfo.Size()
		}

		return err
	})

	return size, err
}

func convertSizeToGb(data int64) string {
	v := units.NewValue(float64(data), units.Byte)
	c := v.MustConvert(units.Gibibyte)

	opts := units.FmtOptions{
		Label:     true,
		Short:     true,
		Precision: 2,
	}

	return c.Fmt(opts)
}
