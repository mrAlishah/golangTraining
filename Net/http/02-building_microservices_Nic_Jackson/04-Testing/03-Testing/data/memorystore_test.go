package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturns1KittenWhenSearchGarfield(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Garfield")

	assert.Equal(t, 1, len(kittens))
}

func TestReturns0KittenWhenSearchTom(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Tom")

	assert.Equal(t, 0, len(kittens))
}
