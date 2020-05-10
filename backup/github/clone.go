package github

import (
	"fmt"
	"io/ioutil"
	"sync"

	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	go_git_ssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

//Save Does a backup of all repositories of the github organization.
func Save(folder string, user string, githubToken string, githubOrganization string, sshPrivateKey string) {
	wg := &sync.WaitGroup{}

	ghRepos := getAllGithubRepositories(githubToken, githubOrganization)
	repos := transformGhReposToRepos(ghRepos)

	clone(folder, user, repos, sshPrivateKey, wg)

	wg.Wait()
}

func clone(folder string, user string, repos []Repository, sshPrivateKey string, wg *sync.WaitGroup) {
	for _, repo := range repos {
		url := fmt.Sprintf("git@github.com:%s/%s.git", user, repo.Name)
		wg.Add(1)
		go cloneRepository(folder, repo.Name, url, sshPrivateKey, wg)
	}
}

func cloneRepository(folder string, projectName string, url string, sshPrivateKey string, wg *sync.WaitGroup) {
	directory := fmt.Sprintf("%s/%s/", folder, projectName)

	_, err := git.PlainClone(
		directory,
		false,
		&git.CloneOptions{
			URL:  url,
			Auth: getSSHKeyAuth(sshPrivateKey),
		},
	)

	fmt.Printf("Cloning %s in directory %s ... ", projectName, directory)

	if err != nil {
		fmt.Println(fmt.Errorf("repository clone fail because: %s", err))
	}

	fmt.Println("")

	wg.Done()
}

func getSSHKeyAuth(privateSSHKeyFile string) transport.AuthMethod {
	var auth transport.AuthMethod
	sshKey, _ := ioutil.ReadFile(privateSSHKeyFile)
	signer, _ := ssh.ParsePrivateKey([]byte(sshKey))
	auth = &go_git_ssh.PublicKeys{User: "git", Signer: signer}
	return auth
}
