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

		err = checkoutCommitId(repo, targetOid)
		if err != nil {
			panic(err)
		}

		fmt.Printf("I checked out commit %s. Dive in!\n", targetCommitHash)

		return nil
	}

	// default behavior: do a revwalk, find and checkout the oldest commit
	commit, err := RevwalkFromHead(repo)

	commitId := commit.Object.Id()
	if commitId == nil {
		panic(fmt.Sprintf("Error: Commit %v not found.\n", commit.Object))
	}

	err = checkoutCommitId(repo, commitId)
	if err != nil {
		panic(err)
	}

	fmt.Println("I checked out the oldest commit. Dive in!")

	return nil
}

func getOidFromHashString(hash string) (*git.Oid, error) {
	return git.NewOid(hash)
}

func checkoutCommitId(repo *git.Repository, commitId *git.Oid) error {
	err := repo.SetHeadDetached(commitId)
	if err != nil {
		return err
	}

	err = repo.CheckoutHead(nil)
	return err
}
