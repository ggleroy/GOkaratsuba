package main

import (
	"fmt"
	"math"
	"strconv"
)

func karatsuba(x, y int) int {
	if x < 10 || y < 10 {
		return x * y
	} else {
		n := int(math.Max(float64(len(strconv.Itoa(x))), float64(len(strconv.Itoa(y)))))
		half := n / 2
		a := x / int(math.Pow10(half))
		b := x % int(math.Pow10(half))
		c := y / int(math.Pow10(half))
		d := y % int(math.Pow10(half))
		ac := karatsuba(a, c)
		bd := karatsuba(b, d)
		ad_plus_bc := karatsuba(a+b, c+d) - ac - bd
		return ac*int(math.Pow10(2*half)) + ad_plus_bc*int(math.Pow10(half)) + bd
	}
}

func main() {
	var x, y int
	fmt.Println("\nWelcome to Karatsuba algorithm!\n\nType your first integer number.")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println("Error reading first number:\n", err)
		return
	}

	fmt.Println("Now type your second integer number.\n")
	_, err = fmt.Scanf("%d", &y)
	if err != nil {
		fmt.Println("Error reading second number:", err)
		return
	}

	fmt.Printf("The multiplication of %d time %d is %d.\n", x, y, karatsuba(x, y))
}
