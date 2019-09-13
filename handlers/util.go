package handlers

import (
	"os"

	git "gopkg.in/libgit2/git2go.v24"
)

func getRepoFromPwd() (*git.Repository, error) {
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

func revwalkFromHead(repo *git.Repository) (*git.Commit, error) {
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

func revwalkNext(repo *git.Repository, currentheadId *git.Oid) (*git.Commit, error) {
	walk, err := repo.Walk()
	if err != nil {
		return nil, err
	}
	defer walk.Free()

	var nextCommit *git.Commit

	walk.Sorting(git.SortTopological | git.SortReverse)

	// Normally you would push the current head and traverse toward ancestors
	// and you would hide the commit you have visited, along with its ancestors.
	// In reverse topology, the range is defined as in the regular case.
	// Reverse sorting will only affect the direction of traversal --
	// from parents toward children. Since we want a range that starts from
	// the current head and ends in the original head, we will need to push
	// the original and hide the current. The walk will start from the current
	// commit's child.
	originHeadId, err := getOriginHeadId(repo)
	if err != nil {
		return nil, err
	}

	walk.Hide(currentheadId)
	walk.Push(originHeadId)
	iterator := func(commit *git.Commit) bool {
		if commit.Parent(0) != nil && currentheadId.Equal(commit.Parent(0).Object.Id()) {
			nextCommit = commit
			return false
		}
		return true 
	}

	err = walk.Iterate(iterator)
	if err != nil {
		return nil, err
	}

	return nextCommit, nil
}

func getOidFromHashString(hash string) (*git.Oid, error) {
	return git.NewOid(hash)
}

// resetToCommitId performs a hard reset to rewind the working directory to match the target commit
func resetToCommitId(repo *git.Repository, commit *git.Commit) error {
	err := repo.ResetToCommit(commit, git.ResetHard, nil)

	return err
}

// getOriginHead peels the origin/HEAD reference to retrieve a commit id
func getOriginHeadId(repo *git.Repository) (*git.Oid, error) {
	branch, err := repo.LookupBranch("origin/HEAD", git.BranchRemote)
	if err != nil {
		return nil, err
	}

	origHeadObj, err := branch.Reference.Peel(git.ObjectCommit)
	if err != nil {
		return nil, err
	}

	return origHeadObj.Id(), nil
}