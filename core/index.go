package core

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	// Indicates that the entry is a new file.
	IndexEntryFlagNew uint32 = 1 << 0
	// Indicates that the entry has been deleted.
	IndexEntryFlagDeleted uint32 = 1 << 1
	// Indicates that the entry is a symbolic link.
	IndexEntryFlagSymlink uint32 = 1 << 2
	// Indicates that the entry is a submodule.
	IndexEntryFlagSubmodule uint32 = 1 << 3
	// Indicates that the entry has been modified.
	IndexEntryFlagModified uint32 = 1 << 4
	// Indicates that the entry is a vecolink.
	IndexEntryFlagVecolink uint32 = 1 << 5
	// Indicates that the entry is a conflict.
	IndexEntryFlagConflict uint32 = 1 << 6
)

type IndexEntry struct {
	RelativePath string
	BlobHash     string
	Mode         FileMode
	Size         int64
	ModTime      time.Time
	Stage        int
	Flags        uint32
}

type Index struct {
	Entries []*IndexEntry
}

func LoadIndex(repoRoot string) (*Index, error) {
	indexPath := filepath.Join(repoRoot, ".veco", "index")
	file, err := os.Open(indexPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Index{}, nil // Return empty index if file doesn't exist
		}
		return nil, err
	}
	defer file.Close()

	index := &Index{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 5 {
			return nil, fmt.Errorf("invalid index entry: %s", line)
		}

		mode, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid mode in index entry: %s", line)
		}

		size, err := strconv.ParseInt(parts[3], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid size in index entry: %s", line)
		}

		modTime, err := time.Parse(time.RFC3339, parts[4])
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp in index entry: %s", line)
		}

		index.Entries = append(index.Entries, &IndexEntry{
			RelativePath: parts[0],
			BlobHash:     parts[1],
			Mode:         FileMode(mode),
			Size:         size,
			ModTime:      modTime,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return index, nil
}

func (index *Index) AddEntry(relPath, blobHash string, info os.FileInfo) {
	entry := IndexEntry{
		RelativePath: relPath,
		BlobHash:     blobHash,
		Mode:         FileMode(info.Mode()),
		Size:         info.Size(),
		ModTime:      info.ModTime(),
	}
	index.Entries = append(index.Entries, &entry)
}

func (index *Index) SaveIndex(repoRoot string) error {
	indexPath := filepath.Join(repoRoot, ".veco", "index")
	file, err := os.Create(indexPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, entry := range index.Entries {
		line := fmt.Sprintf("%s %s %d %d %s\n", entry.RelativePath, entry.BlobHash, entry.Mode, entry.Size, entry.ModTime.Format(time.RFC3339))
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
