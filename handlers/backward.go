package handlers

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
 *	BackwardHandler checks out the child commit of the current HEAD
 */
func BackwardHandler(c *cli.Context) error {
	repo, err := getRepoFromPwd()
	if err != nil {
		panic(err)
	}

	// get current HEAD
	head, err := repo.Head()
	if err != nil {
		panic(err)
	}

	// get commit id and look up parent
	current, err := repo.LookupCommit(head.Target())
	if err != nil {
		panic(err)
	}
	if current.Parent(0) == nil {
		fmt.Println("You're already at the earliest commit.")
		return nil
	}

	prevCommit, err := revwalkPrev(repo, current.Parent(0).Object.Id())
	if err != nil {
		panic(err)
	}
	
	err = resetToCommitId(repo, prevCommit)
	if err != nil {
		panic(err)
	}

	fmt.Println("Checked out the previous commit.")
	return nil
}