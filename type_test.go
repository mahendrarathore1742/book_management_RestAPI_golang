package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewBookTest(t *testing.T) {
	acc, err := NewBook("1", "Mahendra", "java")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)
}

func NewAuthorTest(t *testing.T) {
	acc, err := NewAuthor("mahendra", "singh")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)
}
