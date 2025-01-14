package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

const (
	// VecoDir     = ".veco"
	// ObjectsDir  = "objects"
	// RefsDir     = "refs"
	// HeadFile    = "HEAD"
	// ConfigFile  = "config"
	// IndexFile   = "index"
	BranchesDir = "refs/heads"
)

// GenerateHash generates a SHA-1 hash based on content and timestamp
func GenerateHash() string {
	hasher := sha1.New()
	timestamp := time.Now().String()
	hasher.Write([]byte(timestamp))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GetUserInfo returns the current user's name and email
func GetUserInfo() string {
	currentUser, err := user.Current()
	if err != nil {
		return "unknown"
	}
	return currentUser.Username
}

// CreateDirIfNotExists creates a directory if it doesn't exist
func CreateDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

// IsDirectory checks if the given path is a directory
func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// GetFileHash generates a hash for a file's contents
func GetFileHash(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	hasher := sha1.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// GetVecoRoot finds the root directory of the veco repository
func GetVecoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, VecoDir)); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
