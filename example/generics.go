package example

func LastInt(s []int) int {
	return s[len(s)-1]
}

func LastString(s []string) string {
	return s[len(s)-1]
}

func Last[T any](s []T) T {
	return s[len(s)-1]
}

func contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
