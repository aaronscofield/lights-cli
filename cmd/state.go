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

// stateCmd represents the state command
var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Get the current state of your light",

	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")
		apiKey := viper.GetString("api-key")
		deviceModel := viper.GetString("device-model")
		deviceId := viper.GetString("device-id")

		return stateAction(apiRoot, apiKey, deviceModel, deviceId)
	},
}

func stateAction(apiRoot string, apiKey string, deviceModel string, deviceId string) error {
	_, err := govee.State(apiRoot, apiKey, deviceModel, deviceId)
	return err
}

func init() {
	rootCmd.AddCommand(stateCmd)
}
