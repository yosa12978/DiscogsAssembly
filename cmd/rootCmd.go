package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getReleaseCmd)
	rootCmd.AddCommand(setTokenCmd)
	rootCmd.AddCommand(downloadReleaseCmd)
	rootCmd.AddCommand(whoamiCmd)
	rootCmd.AddCommand(metadataCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

var rootCmd = &cobra.Command{
	Use:     "discasm",
	Short:   "Discasm is a tool for downloading images from discogs release",
	Version: "0.1-alpha",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}
