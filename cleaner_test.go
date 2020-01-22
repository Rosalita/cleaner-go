package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	/////////////////////////////////////////////
	// 1. Table Tests
	// Don't put the tests declaration for table tests at the bottom of a test function.
	// Do put tests declaration as close to the top of a test function as possible.
	tests := []struct {
		stuff  string
		result string
		err    error
	}{
		{"work", "", errors.New("was not fun")},
		{"fun", "that was fun", nil},
	}

	for _, test := range tests {
		result, err := doMoreStuff(test.stuff)
		assert.Equal(t, test.result, result)
		assert.Equal(t, test.err, err)

	}
}

/////////////////////////////////////////////
// 2. Don't test the standard library, because standard library code is already tested extensively
func TestStandardLibrary(t *testing.T) {
	name := "Rosie"
	result := fmt.Sprintf("%s is awesome", name)
	assert.Equal(t, "Rosie is awesome", result)
}
