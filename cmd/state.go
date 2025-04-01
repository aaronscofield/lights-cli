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

// stateCmd represents the state command
var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "Get the current state of your light",

	RunE: func(cmd *cobra.Command, args []string) error {
		apiRoot := viper.GetString("api-root")
		apiKey := viper.GetString("api-key")
		deviceModel := viper.GetString("device-model")
		deviceId := viper.GetString("device-id")

		return stateAction(os.Stdout, apiRoot, apiKey, deviceModel, deviceId)
	},
}

func stateAction(out io.Writer, apiRoot string, apiKey string, deviceModel string, deviceId string) error {
	_, err := state(apiRoot, apiKey, deviceModel, deviceId)
	return err
}

func init() {
	rootCmd.AddCommand(stateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
