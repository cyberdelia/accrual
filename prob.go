package accrual

import "math"

// CDF returns the cumulative distribution function if the given
// normal function, for the given value.
func cdf(μ, σ, v float64) float64 {
	return ((1.0 / 2.0) * (1 + math.Erf((v-μ)/(σ*math.Sqrt2))))
}

// Phi returns the φ-failure for the given value and distribution.
func phi(v float64, d []int64) float64 {
	μ := mean(d)
	σ := standardDeviation(d)
	return -math.Log10(1 - cdf(μ, σ, v))
}

// Mean returns the mean of the given sample.
func mean(values []int64) float64 {
	if len(values) == 0 {
		return 0.0
	}
	var sum int64
	for _, v := range values {
		sum += v
	}
	return float64(sum) / float64(len(values))
}

// StandardDeviation returns standard deviation of the given sample.
func standardDeviation(v []int64) float64 {
	return math.Sqrt(variance(v))
}

// Variance returns variance if the given sample.
func variance(values []int64) float64 {
	if len(values) == 0 {
		return 0.0
	}
	m := mean(values)
	var sum float64
	for _, v := range values {
		d := float64(v) - m
		sum += d * d
	}
	return sum / float64(len(values))
}
