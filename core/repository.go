package core

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/NahomAnteneh/veco/utils"
)

type FileMode uint16

type Repository struct {
	RootPath string
	VecoPath string
	Branches []*Branch
	Remotes  []*Remote
	Index    *Index
	Head     *Head
	Objects  *ObjectStore
}

type Head struct {
	Ref string
}

func NewRepository(rootPath string) *Repository {
	return &Repository{
		RootPath: rootPath,
		VecoPath: filepath.Join(rootPath, utils.VecoDir),
		Index:    &Index{}, // Initialize the Index field
		Branches: []*Branch{
			{
				Name:           "master",
				Commit:         &Commit{},
				IsDefault:      true,
				RemoteName:     "",
				TrackingBranch: "",
				IsLocal:        true,
				IsRemote:       false,
				IsMerged:       false,
				CreationDate:   time.Now(),
				LastUpdated:    time.Now(),
			},
		},
		Remotes: []*Remote{},
		Head: &Head{
			Ref: "refs/heads/master",
		},
		Objects: nil,
	}
}

func (r *Repository) IsInitialized() bool {
	_, err := os.Stat(r.VecoPath)
	return err == nil
}

func (r *Repository) CreateBranch(branch *Branch) error {
	if branch.Exists(r) {
		return fmt.Errorf("branch %s already exists", branch.Name)
	} else {
		branchPath := filepath.Join(r.VecoPath, utils.RefsDir, utils.HeadsDir, branch.Name)
		if err := os.WriteFile(branchPath, []byte(branch.Commit.Hash()), 0644); err != nil {
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

func (r *Repository) SaveBlob(blob *Blob) error {
	blobPath := filepath.Join(r.VecoPath, utils.ObjectsDir, blob.Hash()[:2], blob.Hash()[2:])
	return os.WriteFile(blobPath, blob.Content, 0644)
}
