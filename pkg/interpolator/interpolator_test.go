package interpolator

import (
	"testing"
)

var (
	val1 = 0.0
	val2 = 64.0
	val3 = 128.0
	val4 = 255.0
	t    = 0.675
)

func rawCubic(val1, val2, val3, val4, t float64) float64 {
	return (1-t)*(1-t)*(1-t)*val1 +
		(1-t)*(1-t)*t*val2 +
		(1-t)*t*t*val3 +
		t*t*t*val4
}

func preparedCubic(val1, val2, val3, val4, t float64) float64 {
	tSqr := t * t
	oneMinusT := 1 - t
	oneMinusTSqr := oneMinusT * oneMinusT

	return oneMinusTSqr*(oneMinusT*val1+t*val2) + tSqr*(oneMinusT*val3+t*val4)
}

func BenchmarkRawCubicCalculus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawCubic(val1, val2, val3, val4, t)
	}
}

func BenchmarkPreparedCubicCalculus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		preparedCubic(val1, val2, val3, val4, t)
	}
}
