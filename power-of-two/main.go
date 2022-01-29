package main

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	return n == 1 || n&(n-1) == 0
}
