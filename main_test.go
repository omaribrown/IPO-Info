package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindIpo(t *testing.T) {
	t.Run("Should find IPO Data", func(t *testing.T) {
		// arrange
		testStock := "a"

		// act
		testResult := findIPO(testStock)

		// assert
		assert.Equal(t, "Stock: A IPO'd on 11 18 1999 at $ 30.6572", testResult)
	})
}
