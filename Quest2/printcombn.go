package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	if n < 1 || n > 9 {
		return
	} else if n == 1 {
		for i := '0'; i <= '9'; i++ {
			z01.PrintRune(i)
			if i != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	} else {
		digits := make([]int, n)
		fmt.Println(digits)
		// init
		// for i := range n {
		// 	digits[i] = i
	}

	// 		for {
	// 			// 1. print
	// 			for i := range n {
	// 				z01.PrintRune(rune(digits[i] + '0'))
	// 			}

	// 			// 2. check last
	// 			if digits[0] == 10-n {
	// 				break
	// 			}

	// 			z01.PrintRune(',')
	// 			z01.PrintRune(' ')

	// 			// 3. generate next
	// 			i := n - 1
	// 			for i >= 0 && digits[i] == 10-n+i {
	// 				i--
	// 			}

	// 			digits[i]++

	//			for j := i + 1; j < n; j++ {
	//				digits[j] = digits[j-1] + 1
	//			}
	//		}
	//	}
	//
	// z01.PrintRune('\n')
}
