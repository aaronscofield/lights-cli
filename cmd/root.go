/*
Copyright © 2025 Aaron Scofield
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lights",
	Short: "A simple CLI to control my smart lights",
}

var cfgFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".lights")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todoClient.yaml)")
	rootCmd.PersistentFlags().String("api-root",
		"https://developer-api.govee.com/v1/devices", "Govee API URL")
	rootCmd.PersistentFlags().String("api-key",
		"", "Govee API Key")
	rootCmd.PersistentFlags().String("device-id",
		"", "Govee Device ID")
	rootCmd.PersistentFlags().String("device-model",
		"", "Govee Device Model")

	// Parse environment vairables to find matching flag values
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("LIGHTS")

	viper.BindPFlag("api-root", rootCmd.PersistentFlags().Lookup("api-root"))
	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("device-id", rootCmd.PersistentFlags().Lookup("device-id"))
	viper.BindPFlag("device-model", rootCmd.PersistentFlags().Lookup("device-model"))
}
