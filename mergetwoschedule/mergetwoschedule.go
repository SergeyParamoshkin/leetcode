package mergetwoschedule

type Point struct {
	Time  int
	Value int
}

func Merge(a []Point, b []Point) []Point {
	currentValueA, currentValueB := 0, 0
	indA, indB := 0, 0
	result := []Point{}

	for indA < len(a) && indB < len(b) {
		if a[indA].Time < b[indB].Time {
			currentValueA = a[indA].Value

			result = append(result, Point{Time: a[indA].Time, Value: currentValueA + currentValueB})
			indA++
		} else {
			currentValueB = b[indB].Value

			result = append(result, Point{Time: b[indB].Time, Value: currentValueA + currentValueB})
			indB++
		}
	}

	for indA < len(a) {
		currentValueA = a[indA].Value

		result = append(result, Point{Time: a[indA].Time, Value: currentValueA + currentValueB})
		indA++
	}

	for indB < len(b) {
		currentValueB = b[indB].Value

		result = append(result, Point{Time: b[indB].Time, Value: currentValueA + currentValueB})
		indB++
	}

	return result
}
