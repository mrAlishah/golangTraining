package data

// Store is an interface used for interacting with the backend datastore
// The reason: we need to create a interface at here because it maybe use to mock testing
type Store interface {
	Search(name string) []Kitten
}
