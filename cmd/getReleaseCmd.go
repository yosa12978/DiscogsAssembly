package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yosa12978/DiscogsAssembly/helpers"
	"github.com/yosa12978/DiscogsAssembly/models"
)

var getReleaseCmd = &cobra.Command{
	Use:     "release",
	Short:   "parse release by it's discogs id",
	Example: "discasm release {release_id}",
	Run:     getRelease,
}

func getRelease(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		return
	}

	token := viper.Get("discogs.token")
	release_url := fmt.Sprintf("https://api.discogs.com/releases/%s?token=%s", args[0], token)

	s, err := helpers.HttpGet(release_url)
	if err != nil {
		fmt.Printf("Error fetching release: %s\n", err.Error())
		return
	}
	var release models.Release
	err = json.Unmarshal([]byte(s), &release)
	if err != nil {
		fmt.Printf("Error fetching release: %s\n", err.Error())
		return
	}

	var artists []string
	for i := 0; i < len(release.Artists); i++ {
		artists = append(artists, release.Artists[i].Name)
	}

	fmt.Println("Id: ", release.Id)
	fmt.Println("URL: ", release.Uri)
	fmt.Println("Title: ", release.Title)
	fmt.Println("Artists: ", artists)
	fmt.Println("Year: ", release.Year)
	fmt.Println("Country: ", release.Country)
}
