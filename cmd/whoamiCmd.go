package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yosa12978/DiscogsAssembly/repos"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display current user information",
	Run:   whoami,
}

func whoami(cmd *cobra.Command, args []string) {
	discogsRepo := repos.NewDiscogsRepo()
	user, err := discogsRepo.GetCurrentUser()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Discogs Id: %d\n", user.Id)
	fmt.Printf("Username: %s\n", user.Username)
}
