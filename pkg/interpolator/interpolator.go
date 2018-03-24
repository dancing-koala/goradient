package interpolator

type Interpolator interface {
	Interpolate(input float64)
}

type linearInterpolator struct {
	src  float64
	diff float64
}

func (l *linearInterpolator) Interpolate(input float64) float64 {
	input = clamp(input, 0.0, 1.0)

	return l.src + l.diff*input
}

func NewLinearInterpolator(src, dst float64) *linearInterpolator {
	return &linearInterpolator{
		src:  src,
		diff: dst - src,
	}
}

func clamp(in, min, max float64) float64 {
	if in < min {
		return min
	} else if in > max {
		return max
	}

	return in
}
