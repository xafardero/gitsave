package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/xafardero/gitsave/backup/github"
)

func main() {
	folder := flag.String("folder", "/tmp/gitsave", "Folder to clone repositorires. Default: /tmp/gitsave")
	sshPrivateKey := flag.String("key", "", "Private key used to clone repositories. Default: $HOME/.ssh/id_rsa")

	githubToken := flag.String("gh-token", "", "Github token ID.")
	githubOrganization := flag.String("org", "", "Github organization.")
	user := flag.String("user", "", "Github user.")

	flag.Parse()

	if *githubOrganization == "" {
		fmt.Println("the flags gh-token should not be empty")
		os.Exit(1)
	}

	if *githubToken == "" {
		fmt.Println("the flags gh-token should not be empty")
		os.Exit(1)
	}

	if *folder == "" {
		fmt.Println("the flags folder should not be empty")
		os.Exit(1)
	}

	if *sshPrivateKey == "" {
		homeFoler, _ := homedir.Dir()

		*sshPrivateKey = fmt.Sprintf("%s/.ssh/id_rsa", homeFoler)
	}

	if *user == "" {
		user = githubOrganization
	}

	github.Save(*folder, *user, *githubToken, *githubOrganization, *sshPrivateKey)
}
