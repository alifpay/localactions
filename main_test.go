package main

import (
	"testing"

	"gotest.tools/assert"
)

func FuzzAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int, j int) {
		got := add(i, j)
		assert.Equal(t, i+j, got, "the values i and j should be summed together properly")
	})
}
