package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setTokenCmd = &cobra.Command{
	Use:     "token",
	Short:   "Update token in configuration",
	Example: "discasm token {token}",
	Run:     setToken,
}

func setToken(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		return
	}
	viper.Set("discogs.token", args[0])
	viper.WriteConfig()
	fmt.Println("discogs api token updated")
}
