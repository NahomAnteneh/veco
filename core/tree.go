package core

type Tree struct {
	Entries []*TreeEntry
}

type TreeEntry struct {
	Mode FileMode
	Type string
	Hash string
	Path string // maybe a file name
}

func (t *Tree) ObjectType() string {
	return "tree"
}

func (t *Tree) Hash() string {
	// TODO: implement hashing algorithm for tree
	return ""
}

func (t *Tree) Size() int64 {
	// calculate the size of the tree
	return 0
}
