package wordz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomNumbersImpl_NextInt(t *testing.T) {
	type args struct {
		upperBound int
	}
	tests := []struct {
		name      string
		args      args
		lowRange  int
		highRange int
	}{
		{
			name: "Happy Path",
			args: args{
				upperBound: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ra := RandomNumbersImpl{}
			assert.Condition(t, func() bool {
				n := ra.NextInt(tt.args.upperBound)
				return 0 <= n && n <= tt.args.upperBound
			}, "Random number is out of range")
		})
	}
}
