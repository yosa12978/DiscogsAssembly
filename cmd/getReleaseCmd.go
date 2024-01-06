package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yosa12978/DiscogsAssembly/models"
	"github.com/yosa12978/DiscogsAssembly/repos"
)

var getReleaseCmd = &cobra.Command{
	Use:     "release",
	Short:   "Fetch release by it's discogs id",
	Example: "discasm release {release_id}",
	Run:     getRelease,
}

func getRelease(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		return
	}

	repo := repos.NewDiscogsRepo()
	release, err := repo.GetRelease(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dto := models.ToMetadata(release)

	fmt.Printf("Id: %d\n", dto.Id)
	fmt.Printf("Discogs Link: %s\n", dto.Uri)
	fmt.Printf("Title: %s\n", dto.Title)
	fmt.Printf("Artists: %s\n", strings.Join(dto.Artists, ", "))
	fmt.Printf("Genres: %s\n", strings.Join(dto.Genres, ", "))
	fmt.Printf("Styles: %s\n", strings.Join(dto.Styles, ", "))
	fmt.Printf("Year: %d\n", dto.Year)
	fmt.Printf("Country: %s\n", dto.Country)
}
