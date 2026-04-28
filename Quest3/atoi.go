package piscine

func Atoi(s string) int {	
	result := 0
	sign := 1
	i := 0

	switch s[0] {
	case '-':
		sign = -1
		i++
	case '+':
		i++
	}

	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		} else {
			result = result*10 + int(s[i]-'0')
		}
	}

	return result * sign
}