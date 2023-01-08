package i1346_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/i1346"
)

func TestCheckIfExist(t *testing.T) {
	t.Parallel()

	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "okok", args: args{arr: []int{10, 2, 5, 3}}, want: true},
		// {name: "badbad", args: args{arr: []int{3, 1, 7, 11}}, want: false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := i1346.CheckIfExist(tt.args.arr); got != tt.want {
				t.Errorf("CheckIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
