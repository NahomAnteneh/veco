package core

import "time"

type Branch struct {
	Name           string
	Commit         *Commit
	IsDefault      bool
	RemoteName     string //The name of the remote repository associated with this branch.
	TrackingBranch string //The name of the branch that this branch is tracking (usually in a remote repository).
	IsLocal        bool
	IsRemote       bool
	IsMerged       bool
	CreationDate   time.Time
	LastUpdated    time.Time
}

func (b *Branch) Exists(r *Repository) bool {
	for _, branch := range r.Branches {
		if branch.Name == b.Name {
			return true
		}
	}
	return false
}
