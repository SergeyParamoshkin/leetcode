package fibonacci

func Fib(number int) int {
	const (
		first  = 1
		second = 2
	)

	if number < second {
		return number
	}

	return Fib(number-first) + Fib(number-second)
}
