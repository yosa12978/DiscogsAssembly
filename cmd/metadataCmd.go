package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yosa12978/DiscogsAssembly/repos"
	"github.com/yosa12978/DiscogsAssembly/services"
)

func init() {
	metadataCmd.PersistentFlags().Bool("noimages", false, "use to ignore images in metadata")
	metadataCmd.PersistentFlags().StringP("name", "n", "", "use to specify name of output file")
	metadataCmd.PersistentFlags().StringP("output", "o", ".", "use to specify path of output file")
}

var metadataCmd = &cobra.Command{
	Use:     "metadata",
	Short:   "Download release metadata by it's discogs id",
	Example: "discasm metadata {release_id}",
	Run:     downloadMetadata,
}

func downloadMetadata(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		return
	}

	discogsRepo := repos.NewDiscogsRepo()
	release, err := discogsRepo.GetRelease(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ni, err := cmd.Flags().GetBool("noimages")
	if err != nil {
		cmd.Help()
		return
	}
	// _, err = discogsRepo.GetCurrentUser()
	// isauth := err == nil
	if ni { // || !isauth {
		release.Images = nil
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		cmd.Help()
		return
	}
	if name == "" {
		var artists []string
		for i := 0; i < len(release.Artists); i++ {
			artists = append(artists, release.Artists[i].Name)
		}
		name = fmt.Sprintf("%s - %s", strings.Join(artists, ", "), release.Title)
	}
	path, err := cmd.Flags().GetString("output")
	if err != nil {
		cmd.Help()
		return
	}

	metaServ := services.NewMetadataService()
	if err := metaServ.SaveMetadata(path, name, release); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("This metadata file is already exist. Use --name flag to specify metadata file name")
			return
		}
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("downloaded metadata - %s.json\n", name)
}
