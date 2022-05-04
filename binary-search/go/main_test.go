package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	type args struct {
		nums   []int64
		target int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "ok",
			args: args{
				nums:   []int64{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzSearch(f *testing.F) {
	rand.Seed(time.Now().UnixNano())
	f.Add(10)
	f.Fuzz(func(t *testing.T, n int) {
		n %= 20
		// var vals []int64
		var expect int64
		// for i := 0; i < n; i++ {
		// 	val := rand.Int63() % 1e6
		// 	vals = append(vals, val)
		// 	expect += val
		// }
		expect = 4

		nums := []int64{-1, 0, 3, 5, 9, 12}
		assert.Equal(t, expect, search(nums, 9))
	})

}
