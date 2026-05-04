package main

import (
	"fmt"
	// "github.com/01-edu/z01"
	"piscine"
)

func main() {
	// z01.PrintRune(piscine.LastRune("Hello, World!"))
	// z01.PrintRune(piscine.LastRune("Salut!"))
	// z01.PrintRune(piscine.LastRune("Ola!"))
	// z01.PrintRune('\n')
	// z01.PrintRune(piscine.NRune("Hello!", 3))
	// z01.PrintRune(piscine.NRune("Salut!", 2))
	// z01.PrintRune(piscine.NRune("Bye!", -1))
	// z01.PrintRune(piscine.NRune("Bye!", 5))
	// z01.PrintRune(piscine.NRune("Ola!", 4))
	// z01.PrintRune('\n')
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Salut!", "lut!"))
	fmt.Println(piscine.Compare("Ola!", "Ol"))
}