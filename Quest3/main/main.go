package main

import (
	"fmt"
	"piscine"
)

func main() {
	// n := 0
	// piscine.PointOne(&n)
	// fmt.Println(n)
	// a := 13
	// b := 2
	// var div int
	// var mod int
	// piscine.DivMod(a, b, &div, &mod)
	// fmt.Println(div)
	// fmt.Println(mod)
	// a := 13
	// b := 2
	// piscine.UltimateDivMod(&a, &b)
	// fmt.Println(a)
	// fmt.Println(b)
	// piscine.PrintStr("Hello World!")
	// l := piscine.StrLen("Hello World!")
	// fmt.Println(l)
	// a := 0
	// b := 1
	// piscine.Swap(&a, &b)
	// fmt.Println(a)
	// fmt.Println(b)
	// s := "Hello World!"
	// s = piscine.StrRev(s)
	// fmt.Println(s)
	fmt.Println(piscine.BasicAtoi("12345"))
	fmt.Println(piscine.BasicAtoi("0000000012345"))
	fmt.Println(piscine.BasicAtoi("000000"))
}
