package core

type Object interface {
	ObjectType() string
	Hash() string
	Size() int64
}

type ObjectStore interface {
	Get(hash string) (Object, error)
	Put(obj Object) (string, error)
	Has(hash string) bool
	Delete(hash string) error
}

type InMemoryObjectStore struct {
	objects map[string]Object
}
