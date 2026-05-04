package piscine

func AlphaCount(s string) int {
	var count int
	for _, c := range s {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			count++
		}
	}
	return count
}