package utils

func AddSuffixIfLength(s string, n int, suffix string) string {
	if len(s) == n {
		return s + suffix
	}
	return s
}

func Clamp(n, min, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}
	return n
}
