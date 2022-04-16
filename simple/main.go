package main

import "fmt"

func main() {
	//assignment()
	//ifCondition()
	//switchCondition()
	//fizzBuzz()
	//strings()
	//sprint()
	evenEndedNumber()
}

func assignment() {
	// assignment, mean
	x, y := 1.0, 2.0
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	// this results in a float - the resulting type will follow x and y, not 2
	mean := (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)

}

func ifCondition() {
	x := 1.0
	if x > 5 {
		fmt.Printf("x: %v\n", x)
	}

	if x > 100 {
		fmt.Printf("x > 100\n")
	} else {
		fmt.Printf("x <= 100\n")
	}
}

func switchCondition() {
	x := 1.0
	switch {
	case x > 100:
		fmt.Println("x > 100")
	case x > 10:
		fmt.Println("x > 10")
	default:
		fmt.Println("x < 10")
	}
}

func fizzBuzz() {
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

func strings() {
	book := "The color of magic"
	fmt.Println(book)

	// 18
	fmt.Println(len(book))

	// book[0] = 84 (type uint8)
	// T = 84
	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	// color o
	fmt.Println(book[4:11])

	// color of magic
	fmt.Println(book[4:])

	// The
	fmt.Println(book[:4])

	//
	poem := `
The road goes ever on
Down from the door where it began`
	fmt.Println(poem)
}

func sprint() {
	n := 42
	// convert 42 to string, store in s
	s := fmt.Sprintf("%d", n)
	// s = 42 (type string)
	fmt.Printf("s = %v (type %T)\n", s, s)
	// s = "42" (type string) -> %q includes quotation marks
	fmt.Printf("s = %q (type %T)\n", s, s)
}

func evenEndedNumber() {
	result := 0

	for curr, last := 1000, 9999; curr <= last; curr++ {
		for i := curr; i <= last; i++ {
			num := curr * i
			str := fmt.Sprintf("%d", num)
			if str[0] == str[len(str)-1] {
				result++
			}
		}
	}

	fmt.Println(result)

}
