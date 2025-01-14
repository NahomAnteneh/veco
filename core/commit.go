package core

import "time"

type Signature struct {
	Name  string
	Email string
	Time  time.Time
}

type Commit struct {
	Tree      *Tree
	Parent    []*Commit
	Author    *Signature
	Committer *Signature
	Message   string
	Signed    bool
}

func (c *Commit) ObjectType() string {
	return "commit"
}

func (c *Commit) Hash() string {
	// TODO: implement hashing algorithm for commit
	return ""
}
