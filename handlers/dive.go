package handlers

import (
	"fmt"

	git "gopkg.in/libgit2/git2go.v24"
	"github.com/urfave/cli"
)

/*
 *	DiveHandler checks out oldest commit in the repo
 */

func DiveHandler(c *cli.Context) error {
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

		err = resetToCommitId(repo, targetCommit)
		if err != nil {
			panic(err)
		}

		fmt.Printf("I checked out commit %s. Dive in!\n", targetCommitHash)

		return nil
	}

	// default behavior: do a revwalk, find and checkout the oldest commit
	commit, err := RevwalkFromHead(repo)
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

func getOidFromHashString(hash string) (*git.Oid, error) {
	return git.NewOid(hash)
}

// resetToCommitId performs a hard reset to rewind the working directory to match the target commit
func resetToCommitId(repo *git.Repository, commit *git.Commit) error {
	err := repo.ResetToCommit(commit, git.ResetHard, nil)

	return err
}
