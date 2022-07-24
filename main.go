package main

import (
	"fmt"
	"os"
	"strings"
)

func carry_or_no(value rune) (int, rune) {
	if value > 70 {
		return 1, 48
	}
	if value <= 70 {
		return 0, value
	}
	return 0, 0 // in theory we should never be here
}

func main() {
	var hexstring string
	fmt.Println("give me a hex thing: ")
	fmt.Scanln(&hexstring)
	//hexstring = "0xFFFF"

	// check for the typical hex prefix
	if strings.HasPrefix(hexstring, "0x") {
		//fmt.Println("has prefix")
	} else {
		fmt.Println("bad input")
		os.Exit(3)
	}
	// creates slices of rune...
	a := []rune(hexstring)

	j := len(a) - 1
	last_element := j
	first_element := 2 // 0,1 are the prefix
	reverse_order := ""
	carry_over := 0
	first_element_bit := 0

	for j > 1 {

		if j == last_element {
			a[j] = a[j] + 1
			carry_over, a[j] = carry_or_no(a[j])
		}

		if j != last_element && carry_over == 1 && j != first_element {
			a[j] = a[j] + 1
			carry_over, a[j] = carry_or_no(a[j])
		}

		if j == first_element && carry_over == 1 {
			a[j] = a[j] + 1
			carry_over, a[j] = carry_or_no(a[j])
			first_element_bit = 1
		}

		if j != last_element && carry_over == 0 {
		} // do nothing
		if j == first_element && carry_over == 0 {
		} // do nothing

		reverse_order = string(a[j]) + reverse_order

		if first_element_bit == 1 {
			reverse_order = string(49) + reverse_order
			carry_over = 0
			first_element_bit = 0
		}

		j-- // decrement
	}
	fmt.Println("0x" + reverse_order)
}
