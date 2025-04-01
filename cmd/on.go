/*
Copyright © 2025 Aaron Scofield
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:          "on",
	Short:        "Turn the lights on",
	SilenceUsage: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")
		apiKey := viper.GetString("api-key")
		deviceModel := viper.GetString("device-model")
		deviceId := viper.GetString("device-id")

		return turnAction(os.Stdout, apiRoot, apiKey, deviceModel, deviceId, "on")
	},
}

func turnAction(out io.Writer, apiRoot string, apiKey string, deviceModel string, deviceId string, operation string) error {
	_, err := turn(apiRoot, apiKey, deviceModel, deviceId, operation)
	return err
}

func init() {
	rootCmd.AddCommand(onCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
