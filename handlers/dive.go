package handlers

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
 *	DiveHandler checks out oldest commit in the repo
 */

func DiveHandler(c *cli.Context) error {
	repo, err := getRepoFromPwd()
	if err != nil {
		panic(err)
	}

	// if -c flag is set, try to checkout target commit
	if c.IsSet("commit") {
		targetCommitHash := c.String("commit")

		targetOid, err := getOidFromHashString(targetCommitHash)
		if err != nil {
			panic(err)
		}

		targetCommit, err := repo.LookupCommit(targetOid)
		if err != nil {
			panic(err)
		}

		err = resetToCommitId(repo, targetCommit)
		if err != nil {
			panic(err)
		}

		fmt.Printf("I checked out commit %s. Dive in!\n", targetCommitHash)

		return nil
	}

	// default behavior: do a revwalk, find and checkout the oldest commit
	commit, err := revwalkFromHead(repo)
	if err != nil {
		panic(fmt.Sprintln("Error: Initial commit not found."))
	}

	err = resetToCommitId(repo, commit)
	if err != nil {
		panic(err)
	}

	fmt.Println("I checked out the oldest commit. Dive in!")

	return nil
}
