package bianrysearch_test

import (
	"testing"

	"github.com/SergeyParamoshkin/leetcode/bianrysearch"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "okok",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := bianrysearch.Search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func FuzzSearch(f *testing.F) {
// 	f.Add(10)

// 	f.Fuzz(func(t *testing.T, n int) {
// 		nums := make([]int64, n)
// 		for i := 0; i < n; i++ {
// 			nBig, err := rand.Int(rand.Reader, big.NewInt(27))
// 			if err != nil {
// 				panic(err)
// 			}
// 			nums[i] = nBig.Int64()
// 		}

// 		expect := nums[n-4]
// 		t.Log(expect)
// 		assert.Equal(t, expect, bianrysearch.Search(nums, expect))
// 	})
// }
