package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yosa12978/DiscogsAssembly/repos"
)

var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "displays current user information",
	Run:   whoami,
}

func whoami(cmd *cobra.Command, args []string) {
	discogsRepo := repos.NewDiscogsRepo()
	user, err := discogsRepo.GetCurrentUser()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Id: %v\n", user.Id)
	fmt.Printf("Username: %v\n", user.Username)
	fmt.Printf("URL: %v\n", user.Resource_url)
	fmt.Printf("Consumer Name: %v\n", user.Consumer_name)
}
