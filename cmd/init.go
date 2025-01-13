package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func CmdInit(args []string) error {
	targetDir := "."
	absTargetDir, _ := filepath.Abs(targetDir)
	if len(args) > 0 {
		targetDir = args[0]
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", targetDir, err)
		}
	}

	if err := os.Chdir(targetDir); err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", targetDir, err)
	}

	dirs := []string{
		vecoDir,
		filepath.Join(vecoDir, objectsDir),
		filepath.Join(vecoDir, objectsDir, packDir),
		filepath.Join(vecoDir, refsDir),
		filepath.Join(vecoDir, refsDir, headsDir),
		filepath.Join(vecoDir, refsDir, tagsDir),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}

	repo := core.NewReposiotry(".")
	masterBranch := &core.Branch{
		Name:      "master",
		Commit:    "",
		IsDefault: true,
	}

	if err := repo.UpdateBranch(masterBranch); err != nil {
		return fmt.Errorf("failed to create master branch %v", err)
	}

	head := &core.Head{
		Ref:           "refs/heads/master",
		CurrentCommit: "",
	}

	if err := repo.UpdateHead(head); err != nil {
		return fmt.Errorf("failed to initialize HEAD %v", err)
	}

	files := map[string]string{
		filepath.Join(vecoDir, configFile): "[]",
	}

	for file, content := range files {
		if err := os.WriteFile(file, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %v", file, err)
		}
	}

	fmt.Printf("Initialized empty Veco repository in %s\n", absTargetDir)
	return nil
}
