package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Abs(t *testing.T) {
	type args struct {
		num int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Successful Positive",
			args: args{
				num: 1,
			},
			want: 1,
		},
		{
			name: "Successful Negative",
			args: args{
				num: -1,
			},
			want: 1,
		},
		{
			name: "Successful 0",
			args: args{
				num: -0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Abs(tt.args.num)
			assert.Equal(t, tt.want, got)
		})
	}
}
