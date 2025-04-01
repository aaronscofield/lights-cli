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

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Get the devices associated with your API key",

	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")
		apiKey := viper.GetString("api-key")

		return devicesAction(apiRoot, apiKey)
	},
}

func devicesAction(apiRoot string, apiKey string) error {
	_, err := govee.Devices(apiRoot, apiKey)
	return err
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}
