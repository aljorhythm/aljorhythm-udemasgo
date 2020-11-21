package main

import (
	"fmt"

	"github.com/aljorhythm/udemasgo-sec1-intro/utility"
)

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
