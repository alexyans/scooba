package handlers

import (
	"fmt"
	"os"

	git "gopkg.in/libgit2/git2go.v24"
	"github.com/urfave/cli"
)

/*
 *	DiveHandler checks out oldest commit in the repo
 */

func DiveHandler(c *cli.Context) error {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	repo, err := git.OpenRepository(path)
	if err != nil {
		panic(err)
	}

	walk, err := repo.Walk()
	if err != nil {
		panic(err)
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
		panic(err)
	}

	currentCommitId := currentCommit.Object.Id()
	if currentCommitId == nil {
		panic(fmt.Sprintf("Error: Commit %v not found.\n", currentCommit.Object))
	}

	err = repo.SetHeadDetached(currentCommitId)
	if err != nil {
		panic(err)
	}

	err = repo.CheckoutHead(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("I checked out the oldest commit. Dive in!")

	return nil
}
