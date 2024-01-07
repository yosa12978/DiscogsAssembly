package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yosa12978/DiscogsAssembly/repos"
	"github.com/yosa12978/DiscogsAssembly/services"
)

func init() {
	downloadReleaseCmd.PersistentFlags().StringP("output", "o", ".", "output path")
	downloadReleaseCmd.PersistentFlags().Bool("nometa", false, "download only images")
	downloadReleaseCmd.PersistentFlags().String("metaname", "", "metadata file name")
}

var downloadReleaseCmd = &cobra.Command{
	Use:     "download",
	Short:   "Download release images and metadata by it's discogs id",
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
		fmt.Println("use discasm token {token} to authenticate using discogs api token")
		return
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		cmd.Help()
		return
	}

	// finding release by its id
	release, err := repo.GetRelease(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//download images
	imgService := services.NewImageService()
	if err := imgService.Download(release, output); err != nil {
		fmt.Println(err.Error())
		return
	}

	//download metadata
	nometa, err := cmd.Flags().GetBool("nometa")
	if err != nil {
		fmt.Println(err.Error())
		cmd.Help()
		return
	}
	if !nometa {
		metaname, err := cmd.Flags().GetString("metaname")
		if err != nil {
			fmt.Println(err.Error())
			cmd.Help()
			return
		}
		if metaname == "" {
			metaname = release.Title
		}
		metaServ := services.NewMetadataService()
		path := fmt.Sprintf("%s/%s %d", output, release.Title, release.Id)
		if err := metaServ.SaveMetadata(path, metaname, release); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println("This metadata file is already exist. Use --metaname flag to specify metadata file name")
				return
			}
			fmt.Println(err.Error())
			return
		}
	}
}
