package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NahomAnteneh/veco/core"
	"github.com/NahomAnteneh/veco/utils"
)

func CmdInit(args []string) error {
	targetDir := "."
	if len(args) > 0 {
		targetDir = args[0]
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", targetDir, err)
		}
	}

	absTargetDir, _ := filepath.Abs(targetDir)

	if err := os.Chdir(targetDir); err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", targetDir, err)
	}

	repo := core.NewRepository(".")

	if repo.IsInitialized() {
		fmt.Printf("Reinitialized existing Veco repository in %s\n", absTargetDir)
		return nil
	}

	dirs := []string{
		utils.VecoDir,
		filepath.Join(utils.VecoDir, utils.ObjectsDir),
		filepath.Join(utils.VecoDir, utils.ObjectsDir, utils.PackDir),
		filepath.Join(utils.VecoDir, utils.RefsDir),
		filepath.Join(utils.VecoDir, utils.RefsDir, utils.HeadsDir),
		filepath.Join(utils.VecoDir, utils.RefsDir, utils.TagsDir),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	masterBranch := repo.Branches[0]

	if err := repo.CreateBranch(masterBranch); err != nil {
		return fmt.Errorf("failed to create master branch %v", err)
	}

	head := &core.Head{
		Ref: "refs/heads/master",
	}

	if err := repo.UpdateHead(head); err != nil {
		return fmt.Errorf("failed to initialize HEAD %v", err)
	}

	files := map[string]string{
		filepath.Join(utils.VecoDir, utils.ConfigFile): "{isBare: false}",
	}

	for file, content := range files {
		if err := os.WriteFile(file, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %v", file, err)
		}
	}

	fmt.Printf("Initialized empty Veco repository in %s\n", absTargetDir)
	return nil
}
