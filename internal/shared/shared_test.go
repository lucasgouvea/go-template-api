package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	words := []string{
		"TestCamelCase",
		"lowercase",
		"Class",
		"MyClass",
		"MyC",
		"HTML",
		"PDFLoader",
	}
	expected := [][]string{
		{"Test", "Camel", "Case"},
		{"lowercase"},
		{"Class"},
		{"My", "Class"},
		{"My", "C"},
		{"HTML"},
		{"PDF", "Loader"},
	}
	for i, w := range words {
		slice := SplitCamelCase(w)
		assert.Equal(t, expected[i], slice, "slice should be the same as expected")
	}

}
