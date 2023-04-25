package data

var data = []Kitten{
	Kitten{
		Id:     "1",
		Name:   "Felix",
		Weight: 12.3,
	},
	Kitten{
		Id:     "2",
		Name:   "Fat Freddy's Cat",
		Weight: 20.0,
	},
	Kitten{
		Id:     "3",
		Name:   "Garfield",
		Weight: 35.0,
	},
}

// MemoryStore is a simple in memory datastore that implements Store
// The reason, we need to create this struct is store a list of data Kitten
type MemoryStore struct {
}

// Search returns a slice of Kitten which have a name matching the name in the parameters
// We use pointer receiver to store a slice of Kitten
func (m *MemoryStore) Search(name string) []Kitten {
	// Create a slice of kitten because it don't specific size for array
	// Another way to create a array in Go: var kittens [10]Kitten
	// ==> It will initialize data = zero for all element in this array
	var kittens []Kitten

	// For loop in data, and select only second items in data.
	for _, k := range data {
		if k.Name == name {
			kittens = append(kittens, k)
		}
	}

	return kittens
}
