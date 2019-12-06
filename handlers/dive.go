package handlers

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
 *	DiveHandler checks out oldest commit in the repo
 */

func DiveHandler(c *cli.Context) error {
	summary, err := Dive(c)
	if err != nil {
		return err
	}

	if c.IsSet("commit") {
		fmt.Printf("I checked out the target commit. Dive in!\n")
	} else {
		fmt.Println("I checked out the oldest commit. Dive in!")
	}

	fmt.Println(summary)
	return nil
}

func Dive(c *cli.Context) (*Summary, error) {
	repo, err := GetRepoFromPwd()
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

		err = resetWorktreeToCommit(repo, targetCommit)
		if err != nil {
			panic(err)
		}

		return NewSummary(targetCommit), nil
	}

	// default behavior: do a revwalk, find and checkout the oldest commit
	commit, err := revwalkFromHead(repo)
	if err != nil {
		panic(fmt.Sprintln("Error: Initial commit not found."))
	}

	err = resetWorktreeToCommit(repo, commit)
	if err != nil {
		panic(err)
	}

	return NewSummary(commit), nil
}
