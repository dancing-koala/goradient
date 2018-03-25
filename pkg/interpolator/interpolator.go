package interpolator

type Interpolator interface {
	Interpolate(input float64) float64
}

type linearInterpolator struct {
	src float64
	dst float64
}

func (l *linearInterpolator) Interpolate(t float64) float64 {
	t = clamp(t, 0.0, 1.0)
	return (1.0-t)*l.src + t*l.dst
}

func NewLinear(src, dst float64) *linearInterpolator {
	return &linearInterpolator{
		src: src,
		dst: dst,
	}
}

type quadraticInterpolator struct {
	first  float64
	second float64
	third  float64
}

func (q *quadraticInterpolator) Interpolate(t float64) float64 {
	t = clamp(t, 0.0, 1.0)

	oneMinusT := 1.0 - t

	return oneMinusT*(oneMinusT*q.first+2*t*q.second) + t*t*q.third
}

func NewQuadratic(first, second, third float64) *quadraticInterpolator {
	return &quadraticInterpolator{
		first:  first,
		second: second,
		third:  third,
	}
}

type cubicInterpolator struct {
	first  float64
	second float64
	third  float64
	fourth float64
}

func (c *cubicInterpolator) Interpolate(t float64) float64 {
	t = clamp(t, 0.0, 1.0)
	tSqr := t * t
	oneMinusT := 1 - t
	oneMinusTSqr := oneMinusT * oneMinusT

	return oneMinusTSqr*(oneMinusT*c.first+t*c.second) + tSqr*(oneMinusT*c.third+t*c.fourth)
}

func NewCubic(first, second, third, fourth float64) *cubicInterpolator {
	return &cubicInterpolator{
		first:  first,
		second: second,
		third:  third,
		fourth: fourth,
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
