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

func revwalkFromHead(repo *git.Repository) (*git.Commit, error) {
	walk, err := repo.Walk()
	if err != nil {
		return nil, err
	}
	defer walk.Free()

	var currentCommit *git.Commit

	walk.Sorting(git.SortTopological | git.SortReverse)
	
	originHeadId, err := getOriginHeadId(repo)
	if err != nil {
		return nil, err
	}

	walk.Push(originHeadId)
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

func revwalkNext(repo *git.Repository, currentHeadId *git.Oid) (*git.Commit, error) {
	walk, err := repo.Walk()
	if err != nil {
		return nil, err
	}
	defer walk.Free()

	var nextCommit *git.Commit
	walk.Sorting(git.SortTopological)

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

	// walk.Hide(currentHeadId)
	walk.Push(originHeadId)

	iterator := func(commit *git.Commit) bool {
		if isParent(currentHeadId, commit) {
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

func isParent(ancestorId *git.Oid, currentCommit *git.Commit) bool {
	if currentCommit.Parent(0) == nil && currentCommit.Parent(1) == nil {
		return false
	}

	if currentCommit.Parent(0) != nil && ancestorId.Equal(currentCommit.Parent(0).Object.Id()) {
		return true
	}

	if currentCommit.Parent(1) != nil && ancestorId.Equal(currentCommit.Parent(1).Object.Id()) {
		return true
	}
	
	return false
}

func revwalkPrev(repo *git.Repository, currentHeadId *git.Oid) (*git.Commit, error) {
	walk, err := repo.Walk()
	if err != nil {
		return nil, err
	}
	defer walk.Free()

	var prevCommit *git.Commit
	walk.Sorting(git.SortTopological)

	// Push the current head to traverse toward its ancestors
	walk.PushHead()
	iterator := func(commit *git.Commit) bool {
		if commit.Parent(0) != nil && currentHeadId.Equal(commit.Parent(0).Object.Id()) {
			prevCommit = commit.Parent(0)
			return false
		}
		
		return true
	}

	err = walk.Iterate(iterator)
	if err != nil {
		return nil, err
	}

	return prevCommit, nil
}

func getOidFromHashString(hash string) (*git.Oid, error) {
	return git.NewOid(hash)
}

// resetWorktreeToCommitId performs a hard reset to rewind the working directory to match the target commit
func resetWorktreeToCommit(repo *git.Repository, commit *git.Commit) error {
	err := repo.ResetToCommit(commit, git.ResetHard, nil)

	return err
}

// resetIndexToCommitId performs a mixed reset to rewind the index to match the target commit
func resetIndexToCommit(repo *git.Repository, commit *git.Commit) error {
	err := repo.ResetToCommit(commit, git.ResetMixed, nil)

	return err
}

// resetHeadToCommitId performs a soft rest to set HEAD to the target commit
func resetHeadToCommit(repo *git.Repository, commit *git.Commit) error {
	err := repo.ResetToCommit(commit, git.ResetSoft, nil)

	return err
}

// getOriginHeadId peels the origin/HEAD reference to retrieve a commit id
func getOriginHeadId(repo *git.Repository) (*git.Oid, error) {
	branch, err := repo.LookupBranch("origin/master", git.BranchRemote)
	if err != nil {
		return nil, err
	}

	origHeadObj, err := branch.Reference.Peel(git.ObjectCommit)
	if err != nil {
		return nil, err
	}

	return origHeadObj.Id(), nil
}