package bitutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasUnsignedBitAt(t *testing.T) {
	a := byte(0b0000_0100)
	assert.True(t, HasBitAt(a, 2))
	assert.False(t, HasBitAt(a, 0))

	b := uint8(0b0000_0100)
	assert.True(t, HasBitAt(b, 2))
	assert.False(t, HasBitAt(b, 0))

	c := uint32(0b1000_0100)
	assert.True(t, HasBitAt(c, 2))
	assert.True(t, HasBitAt(c, 7))
	assert.False(t, HasBitAt(c, 0))

	d := uint64(0b0100_0000)
	assert.True(t, HasBitAt(d, 6))
	assert.False(t, HasBitAt(d, 0))
}
