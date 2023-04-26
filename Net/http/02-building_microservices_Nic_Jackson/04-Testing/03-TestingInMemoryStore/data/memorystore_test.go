package data

import (
	"testing"
)

func TestReturns1KittenWhenSearchGarfield(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Garfield")

	if len(kittens) <= 0 {
		t.Error("We expected result")
	}
}

func TestReturns0KittenWhenSearchTom(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Tom")

	if len(kittens) <= 0 {
		t.Error("We expected result")
	}
}
