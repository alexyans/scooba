package handlers

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
 *	ForwardHandler checks out the child commit of the current HEAD
 */
func ForwardHandler(c *cli.Context) error {
	repo, err := getRepoFromPwd()
	if err != nil {
		panic(err)
	}

	// get current HEAD
	head, err := repo.Head()
	if err != nil {
		panic(err)
	}

	nextCommit, err := revwalkNext(repo, head.Target())
	if err != nil {
		panic(err)
	}
	
	err = resetToCommitId(repo, nextCommit)
	if err != nil {
		panic(err)
	}

	fmt.Println("Checked out the next commit.")

	// do a revwalk by specifying the start of the range to be current commit
	// jump by 1, or later X
	// then checkout result

	return nil
}