package util

func Lerp(start, end, t float64) float64 {
	return start*(1-t) + end*t
}
