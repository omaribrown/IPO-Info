package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIpo(t *testing.T) {
	t.Run("Should find IPO Data", func(t *testing.T) {
		// Enter a Stock symbol
		testStock := "a"

		// Assign the function output to a variable
		testResult := findIPO(testStock)

		// Test against expected result
		assert.Equal(t, "11/18/1999", testResult)
	})
}
