package piscine

import "github.com/01-edu/z01"

// func PrintComb2() {
// 	for a := 0; a <= 98; a++ {
// 		for b := a+1; b <= 99; b++ {
// 			z01.PrintRune(rune(a/10) + '0')
// 			z01.PrintRune(rune(a%10) + '0')
// 			z01.PrintRune(' ')
// 			z01.PrintRune(rune(b/10) + '0')
// 			z01.PrintRune(rune(b%10) + '0')
// 			if !(a==98 && b==99) {
// 				z01.PrintRune(',')
// 				z01.PrintRune(' ')
// 			}
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := a; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					if c > a || (a == c && d > b) {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(c)
						z01.PrintRune(d)

						if !(a == '9' && b == '8' && c == '9' && d == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
