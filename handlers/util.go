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

func RevwalkFromHead(repo *git.Repository) (*git.Commit, error) {
	walk, err := repo.Walk()
	if err != nil {
		return nil, err
	}
	defer walk.Free()

	var currentCommit *git.Commit

	walk.Sorting(git.SortTime | git.SortReverse)

	walk.PushHead()
	iterator := func(commit *git.Commit) bool {
		currentCommit = commit
		return false
	}

	err = walk.Iterate(iterator)
	if err != nil {
		return nil, err
	}

	return currentCommit, nil
}
