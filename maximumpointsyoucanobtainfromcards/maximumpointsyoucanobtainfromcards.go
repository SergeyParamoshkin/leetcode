package maximumpointsyoucanobtainfromcards

func MaxScore(cardPoints []int, k int) int {
	n := len(cardPoints) - 1

	l, r := 0, n
	maxSum := 0

	for r > n-k {
		maxSum += cardPoints[r]
		r--
	}

	tempSum := maxSum

	for r < len(cardPoints)-1 {
		r++
		tempSum -= cardPoints[r]
		tempSum += cardPoints[l]
		l++

		if tempSum > maxSum {
			maxSum = tempSum
		}
	}

	return maxSum
}
