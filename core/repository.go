package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/NahomAnteneh/veco/utils"
)

type Repository struct {
	RootPath string
	VecoPath string
}

type Commit struct {
	CommitHash string
	Tree       Tree
	Message    string
	Parent     string
	Timestamp  time.Time
	Author     string
	Commiter   string
}

type Tree struct {
	Mode     int
	Type     string
	Hash     string
	FileName string
}

type Branch struct {
	Name      string
	Commit    string
	IsDefault bool
}

type Head struct {
	Ref string
}

func NewRepository(rootPath string) *Repository {
	return &Repository{
		RootPath: rootPath,
		VecoPath: filepath.Join(rootPath, utils.VecoDir),
	}
}

func (r *Repository) IsInitialized() bool {
	_, err := os.Stat(r.VecoPath)
	return err == nil
}

func (r *Repository) CreateBranch(branch *Branch) error {
	if branch.IsDefault {
		if err := r.UpdateHead(&Head{Ref: fmt.Sprintf("refs/heads/%s", branch.Name)}); err != nil {
			return fmt.Errorf("failed to create branch %s: %v", branch.Name, err)
		}
	} else {
		branchPath := filepath.Join(r.VecoPath, utils.RefsDir, utils.HeadsDir, branch.Name)
		if err := os.WriteFile(branchPath, []byte(branch.Commit), 0644); err != nil {
			return fmt.Errorf("failed to create branch %s: %v", branch.Name, err)
		}
	}
	return nil
}

func (r *Repository) UpdateHead(head *Head) error {
	headPath := filepath.Join(r.VecoPath, utils.HeadFile)
	if err := os.WriteFile(headPath, []byte(head.Ref), 0644); err != nil {
		return fmt.Errorf("failed to update HEAD: %v", err)
	}
	return nil
}
