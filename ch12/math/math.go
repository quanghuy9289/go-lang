package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}

	total := float64(0)
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// Finds the minimum values in a slice of float64
func Min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	min := xs[0]
	for _, v := range xs {
		if v < min {
			min = v
		}
	}
	return min
}

// Finds the maximum values in a slice of float64
func Max(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	max := xs[0]
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return max
}
