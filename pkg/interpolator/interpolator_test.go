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
	return (1.0-t)*(1.0-t)*(1.0-t)*val1 +
		(1.0-t)*(1.0-t)*t*val2 +
		(1.0-t)*t*t*val3 +
		t*t*t*val4
}

func preparedCubic(val1, val2, val3, val4, t float64) float64 {
	tSqr := t * t
	d := 1.0 - t
	dSqr := d * d

	return dSqr*(d*val1+t*val2) + tSqr*(d*val3+t*val4)
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
