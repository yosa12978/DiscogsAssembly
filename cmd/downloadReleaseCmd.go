package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yosa12978/DiscogsAssembly/services"
)

func init() {
	downloadReleaseCmd.PersistentFlags().StringP("output", "o", ".", "output path")
}

var downloadReleaseCmd = &cobra.Command{
	Use:     "download",
	Short:   "downloads release images and metadata by it's discogs id",
	Example: "discasm download {release_id} (flags)",
	Run:     downloadRelease,
}

func downloadRelease(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		return
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		cmd.Help()
		return
	}

	token := viper.Get("discogs.token")
	release_url := fmt.Sprintf("https://api.discogs.com/releases/%s?token=%s", args[0], token)

	imgService := services.NewImageService()
	if err := imgService.Download(release_url, output); err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
}
