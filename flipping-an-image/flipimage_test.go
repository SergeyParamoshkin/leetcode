package flippinganimage_test

import (
	"reflect"
	"testing"

	flippinganimage "github.com/SergeyParamoshkin/leetcode/flipping-an-image"
)

func Test_flipAndInvertImage(t *testing.T) {
	type args struct {
		image [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "ok",
			args: args{image: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}},
			want: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flippinganimage.FlipAndInvertImage(tt.args.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
