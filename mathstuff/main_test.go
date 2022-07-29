package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//run this test!
func TestAddOrSubtract(t *testing.T) {
	a := assert.New(t)
	result := addOrSubtract("add", 1, 2)
	// This should pass
	a.Equal(3, result)

	// This should fail. How could we make it pass?
	result = addOrSubtract("subtract", 5, 4)
	a.Equal(4, result)

	// TODO : Write a third passing case
}
