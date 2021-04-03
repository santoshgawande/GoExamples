package main

import "fmt"

func main() {
	// First Type For
	i := 1
	for i <= 3 {
		// fmt.Print(i)  -> 123 output on one line
		fmt.Println(i) // -> output on newline
		i = i + 1
	}
	// Second
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// Third It execute and then break loop
	for {
		fmt.Println("loop")
		break
	}

	// Fourth I condition true then continue loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
