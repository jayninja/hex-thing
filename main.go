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
	//hexstring = "0xA499F"

	// check for the typical hex prefix
	if strings.HasPrefix(hexstring, "0x") {
		// do nothing, input is valid
	} else {
		fmt.Println("bad input")
		os.Exit(3)
	}

	// Uppercase the entire string.
	hexstring = strings.ToUpper(hexstring)

	// creates slices of rune...
	a := []rune(hexstring)

	// quickly validate the rest of our hex characters
	for iterator := len(hexstring) - 1; iterator > 2; iterator-- {
		if a[iterator] > 70 || a[iterator] < 48 || (a[iterator] < 65 && a[iterator] > 57) {
			fmt.Println("Valid hex characters only please")
			os.Exit(3)
		}
	}

	j := len(a) - 1 // because our index starts at 0, and length at 1
	last_element := j
	first_element := 2     // 0,1 are the prefix so the first one we care about is 2
	first_element_bit := 0 // used to tell when the first element has a carry_over
	carry_over := 0        // decide if we need to carry over
	reverse_order := ""    // string to build our result

	// more than 1 and we are at our first_element
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
		// gross hack to skip over these ascii values.
		if (a[j] > 57) && (a[j] < 64) {
			a[j] = a[j] + 7
		}

		/*
			if j != last_element && carry_over == 0 {
			} // do nothing
			if j == first_element && carry_over == 0 {
			} // do nothing
		*/

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
