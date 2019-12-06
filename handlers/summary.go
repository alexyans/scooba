package handlers

import (
	"fmt"

	git "gopkg.in/libgit2/git2go.v24"
)

type Changeset struct {
	Adds *int64
	Deletes *int64
	Renames *int64
}

type Summary struct {
	Commit *git.Commit
	Hash *string
	Message *string
	LongMessage *string
	Description *string
	Changeset *Changeset
	PullRequests []string
	Issues []string
	Skipped []Summary
}

func NewSummary(commit *git.Commit) *Summary {
	hash := commit.Object.Id().String()
	message := commit.Summary()
	
	return &Summary{
		Commit: commit,
		Hash: &hash,
		Message: &message,
	}	
}

func (s Summary) String() string {
	var out string

	if s.Hash != nil {
		out += fmt.Sprintf("Commit: %s\n", *s.Hash)
	}
	if s.Message != nil {
		out += fmt.Sprintf("Message: %s\n", *s.Message)
	}

	return out
}