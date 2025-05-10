package valutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected int
		wantErr  bool
	}{
		{"int", 42, 42, false},
		{"float64", 3.14, 3, false},
		{"float64", 3.56, 3, false},
		{"true bool", true, 1, false},
		{"valid string", "123", 123, false},
		{"valid float string", "123.456", 123, false},
		{"invalid string", "abc", 0, true},
		{"unsupported type", []int{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MustInt(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				zero := Int(tt.input)
				assert.Zero(t, zero)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
