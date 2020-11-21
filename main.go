package main

import (
	"fmt"

	"github.com/aljorhythm/aljorhythm-udemasgo/utility"
)

// demonstrates higher order function, function closures
func addOrSubtract() {
	operations := map[string](func(int, int) int){"+": utility.Add, "-": utility.Minus}
	x, y := 6, 7
	for name, operation := range operations {
		fmt.Println(x, name, y, "=", operation(x, y))
	}
}

// api call
func printTheySaidSo(endSignal chan bool) {
	go func() {
		bytes, err := utility.Theysaidso()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(*bytes))
		}
		endSignal <- true
	}()
}

func changeX(x *int) {
	*x = 2
}

const (
	one = 1
	two = 2
)

func pointers() {
	x := one
	xPointer := &x
	fmt.Println("old x", x)
	changeX(xPointer)
	fmt.Println("new x", x)
}

func main() {
	fmt.Println(utility.GetHello())
	endSignal := make(chan bool)
	printTheySaidSo(endSignal)
	addOrSubtract()
	pointers()
	<-endSignal
	fmt.Println("end")
}
