/*
Copyright © 2025 Aaron Scofield
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// onCmd represents the on command
var offCmd = &cobra.Command{
	Use:          "off",
	Short:        "Turn the lights off",
	SilenceUsage: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")
		apiKey := viper.GetString("api-key")
		deviceModel := viper.GetString("device-model")
		deviceId := viper.GetString("device-id")

		return turnAction(apiRoot, apiKey, deviceModel, deviceId, "off")
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}
