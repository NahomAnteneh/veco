package core

type Blob struct {
	Content []byte
}

func (b *Blob) ObjectType() string {
	return "blob"
}

func (b *Blob) Hash() string {
	return CreateBlob(b.Content)
}

func (b *Blob) Size() int64 {
	return int64(len(b.Content))
}

func CreateBlob(content []byte) string {
	return ""
}
