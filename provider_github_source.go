package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/libgit2/git2go"
)

// Sorry for the inverted name versus the filename—one is better in the language,
// the other is nice for viewing with `ls` :-/
type githubSourceProvider struct {
	c config
}

func (p githubSourceProvider) AddRoutes(r map[string][]route, c config) map[string][]route {
	proceed, err := p.shouldAddRoutes(c)
	if err != nil {
		fmt.Println("Config is invalid for githubSourceProvider. " + err.Error())
	}
	if proceed {
		r["init"] = append(r["init"], p.init)
	}
	p.c = c
	return r
}

func (p githubSourceProvider) shouldAddRoutes(c config) (bool, error) {
	useThis := contains(c.Use, "github_source")
	if !useThis {
		return false, nil
	}
	return p.configIsValid(c)
}

func (p githubSourceProvider) configIsValid(c config) (bool, error) {

	return true, nil
}

func (p *githubSourceProvider) init(a []string) error {
	fmt.Printf("%+v\n", p.c.GithubSource.Locations)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	for local, remote := range p.c.GithubSource.Locations {
		path := string(append([]rune(wd), filepath.Separator)) + local
		fmt.Println(path + " " + remote)

		authCallbacks := git.RemoteCallbacks{
			CredentialsCallback:      p.credentialsCallback,
			CertificateCheckCallback: p.certificateCheckCallback,
		}

		fetchOpts := git.FetchOptions{
			RemoteCallbacks: authCallbacks,
		}

		cloneOpts := git.CloneOptions{
			FetchOptions: &fetchOpts,
		}
		_, err := git.Clone(remote, path, &cloneOpts)
		if err != nil {
			fmt.Println(err.Error())
			//return err
		}
	}
	return nil
}

func (p *githubSourceProvider) certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
	return 0
}

func (p *githubSourceProvider) credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
	ret, cred := git.NewCredSshKeyFromAgent(username)
	return git.ErrorCode(ret), &cred
}
