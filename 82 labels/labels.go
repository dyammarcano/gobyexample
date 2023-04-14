package main

import "fmt"

func main() {
	// Using a label to break out of nested loops
outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				// Break out of both loops
				break outerLoop
			}
			fmt.Printf("%d%d\n", i, j)
		}
	}
}
