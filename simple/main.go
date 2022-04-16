package main

import "fmt"

func main() {
	// assignment, mean
	x, y := 1.0, 2.0
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	// this results in a float - the resulting type will follow x and y, not 2
	mean := (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)

	// simple
	if x > 5 {
		fmt.Printf("x: %v\n", x)
	}

	if x > 100 {
		fmt.Printf("x > 100\n")
	} else {
		fmt.Printf("x <= 100\n")
	}

	// switch condition
	switch {
	case x > 100:
		fmt.Println("x > 100")
	case x > 10:
		fmt.Println("x > 10")
	default:
		fmt.Println("x < 10")
	}

	// fizzbuzz
	curr := 1
	last := 20
	for curr != last {
		mod3 := curr%3 == 0
		mod5 := curr%5 == 0
		if mod3 {
			fmt.Println("fizz")
		}
		if mod5 {
			fmt.Println("buzz")
		}
		if !mod3 && !mod5 {
			fmt.Println(curr)
		}
		curr++
	}
}
