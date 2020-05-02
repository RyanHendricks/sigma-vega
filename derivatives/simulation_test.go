package derivatives

import (
	"testing"
)

func Test_calcGreeks1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calcGreeks()
		})
	}
}
