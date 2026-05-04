package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		for i := power; i > 0; i-- {
			result *= nb
		}
	}
	return result
}