package core

import "time"

type Repository struct {
	RootPath string
	VecoPath string
}

type Commit struct {
	ID        string    `json:"id"`
	Tree      string    `json:"tree"`
	Message   string    `json:"message"`
	Parent    string    `json:"parent"`
	Parents   []string  `json:"parents,omitempty"` // For merge commits
	Timestamp time.Time `json:"timestamp"`
	Author    string    `json:"author"`
	Commiter  string    `json:"commiter"`
}

type Head struct {
	Ref           string `json:"ref"`            // e.g., "refs/heads/master"
	CurrentCommit string `json:"current_commit"` // Current commit ID
}
