package piscine

func IterativeFactorial(nb int) int {
	if nb < 1 || nb > 20 {
		return 0
	}
	factorial := 1
	for i := nb; i > 1; i-- {
		factorial *= i
	}
	return factorial
}