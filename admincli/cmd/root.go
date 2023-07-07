package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "admincli",
	Short: "Maintenance tools",
	Long:  `The app gives you simple access to maintenance tools.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	viper.SetDefault("project_dir", ".")

    viper.SetConfigName("admincli")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("/etc/admincli/")
    viper.AddConfigPath("$HOME/.admincli/")
    viper.AddConfigPath("./")

    viper.SetEnvPrefix("admincli")
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
