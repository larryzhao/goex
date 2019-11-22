package mathx

func MaxInt(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func MinInt(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

func MaxInt64(a int64, b int64) int64 {
	if a >= b {
		return a
	}

	return b
}

func MaxFloat64(a float64, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}
