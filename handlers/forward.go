package handlers

import (
	"fmt"

	"github.com/urfave/cli"
)

/*
 *	ForwardHandler checks out the child commit of the current HEAD
 */
func ForwardHandler(c *cli.Context) error {
	summary, err := Forward(c)
	if err != nil {
		return err
	}

	fmt.Println("Checked out the next commit in master.")
	fmt.Println(summary)
	return nil
}

func Forward(c *cli.Context) (*Summary, error) {
	repo, err := GetRepoFromPwd()
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
	
	// update working tree to introduce next batch of changes
	err = resetWorktreeToCommit(repo, nextCommit)
	if err != nil {
		panic(err)
	}

	// rewind index to last known state
	// the introduced changed inbetween are in the working tree but not in the index
	// and can be diffed for inspection
	oldHead, err := repo.LookupCommit(head.Target())
	if err != nil {
		panic(err)
	}
	err = resetIndexToCommit(repo, oldHead)
	if err != nil {
		panic(err)
	}

	// rewinding the index also moved HEAD
	// set it back to the newly visited commit ID so navigation can resume from that point
	err = resetHeadToCommit(repo, nextCommit)
	if err != nil {
		panic(err)
	}

	// do a revwalk by specifying the start of the range to be current commit
	// jump by 1, or later X
	// then checkout result

	return NewSummary(nextCommit), nil
}