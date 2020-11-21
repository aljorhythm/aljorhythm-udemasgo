package main

import (
	"fmt"

	"github.com/aljorhythm/aljorhythm-udemasgo/utility"
)

func closure() {

}

func addOrSubtract() {
	operations := map[string](func(int, int) int){"+": utility.Add, "-": utility.Minus}
	x, y := 6, 7
	for name, operation := range operations {
		fmt.Println(x, name, y, "=", operation(x, y))
	}
}

func printTheySaidSo() {
	bytes, err := utility.Theysaidso()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(*bytes))
	}
}

func main() {
	fmt.Println(utility.GetHello())
	printTheySaidSo()
}
