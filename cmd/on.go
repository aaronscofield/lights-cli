/*
Copyright © 2025 Aaron Scofield
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"lights/internal/govee"

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

		return turnAction(apiRoot, apiKey, deviceModel, deviceId, "on")
	},
}

func turnAction(apiRoot string, apiKey string, deviceModel string, deviceId string, operation string) error {
	_, err := govee.Turn(apiRoot, apiKey, deviceModel, deviceId, operation)
	return err
}

func init() {
	rootCmd.AddCommand(onCmd)
}
