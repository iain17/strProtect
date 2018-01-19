package cmd

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestSimple(t * testing.T) {
	expected := 10

	a := expected ^ 2
	b := a ^ 2
	assert.Equal(t, expected, b)
}

func TestGetBytesShort(t *testing.T) {
	msg := "o"
	protected := protect(msg)
	assert.Equal(t, string(unProtect(protected)), msg)
}

func TestGetBytesMessage(t *testing.T) {
	msg := "Hey this works"
	protected := protect(msg)
	assert.Equal(t, string(unProtect(protected)), msg)
}