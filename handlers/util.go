package handlers

import (
	"os"

	git "gopkg.in/libgit2/git2go.v24"
)

func GetRepoFromPwd() (*git.Repository, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	
	repo, err := git.OpenRepository(path)
	if err != nil {
		return nil, err
	}

	return repo, nil
}