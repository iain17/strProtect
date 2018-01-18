package cmd

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestGetBytes(t *testing.T) {
	msg := "This works"
	protected := getBytes(msg)
	assert.Equal(t, string(reverseBytes(string(protected))), msg)
}