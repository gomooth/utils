package valutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustBool(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected bool
		wantErr  bool
	}{
		{"bool true", true, true, false},
		{"bool false", false, false, false},
		{"string true", "true", true, false},
		{"string yes", "yes", true, false},
		{"string false", "false", false, false},
		{"string no", "no", false, false},
		{"empty string", "", false, false},
		{"number string 1", "1", true, false},
		{"number string 0", "0", false, false},
		{"float string 1.22", "1.22", true, false},
		{"float string 1.22 with space", " 1.22 ", true, false},
		{"float string 0.1", "0.1", true, false},
		{"float string -0.1", "-0.1", false, false},
		{"float string -0.1 with space", " -0.1", false, false},
		{"invalid string", "maybe", false, true},
		{"positive int", 42, true, false},
		{"negative int", -1, false, false},
		{"positive float", 3.14, true, false},
		{"negative float", -0.5, false, false},
		{"unsupported type", []int{}, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MustBool(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.False(t, Bool(tt.input))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
