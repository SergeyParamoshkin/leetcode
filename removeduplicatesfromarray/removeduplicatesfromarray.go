package removeduplicatesfromarray

import (
	"fmt"
)

func RemoveDuplicates(m []int) int {

	o := make([]int, len(m))

	for i := 0; i <= len(m)-2; i++ {
		if m[i] == m[i+1] {
			buff := m[i]
			for m[i+1] == buff {
				i++
			}
		}
	}

	fmt.Println(len(m), m, o)

	return len(m)
}
