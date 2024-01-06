package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yosa12978/DiscogsAssembly/repos"
	"github.com/yosa12978/DiscogsAssembly/services"
)

func init() {
	downloadReleaseCmd.PersistentFlags().StringP("output", "o", ".", "output path")
}

var downloadReleaseCmd = &cobra.Command{
	Use:     "download",
	Short:   "Downloads release images and metadata by it's discogs id",
	Example: "discasm download {release_id} (flags)",
	Run:     downloadRelease,
}

func downloadRelease(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		return
	}

	repo := repos.NewDiscogsRepo()
	if _, err := repo.GetCurrentUser(); err != nil {
		fmt.Println("you unable to download images if you are unauthorized")
		fmt.Println("use discasm token {token} to authenticate via discogs api token")
		return
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		cmd.Help()
		return
	}

	imgService := services.NewImageService()
	if err := imgService.Download(args[0], output); err != nil {
		fmt.Println(err.Error())
		return
	}
}
