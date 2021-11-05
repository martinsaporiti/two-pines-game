package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNormalTray(t *testing.T) {
	// when
	try := NewTry(5)

	// then
	assert.Equal(t, 5, try.KnockedDownPins)
	assert.False(t, try.IsFoul())
	assert.True(t, try.hasValue)
}

func TestNewFoulTray(t *testing.T) {
	// when
	try := NewTry(-1)

	// then
	assert.Equal(t, 0, try.KnockedDownPins)
	assert.True(t, try.IsFoul())
	assert.True(t, try.hasValue)
}
