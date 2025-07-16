package main

import (
	"fmt"
	"os"
)

func main() {
	var lim int
	if _, err := fmt.Fscan(os.Stdin, &lim); err != nil {
		fmt.Fprintln(os.Stderr, "Error input", err)
		return
	}
	for i := 0; i < lim; i++ {
		var str string = ""
		if (i+1)%3 == 0 {
			str += "Fizz"
		}
		if (i+1)%5 == 0 {
			str += "Buzz"
		}
		if str != "" {
			fmt.Print(str)
		} else {
			fmt.Print(i + 1)
		}
		if i < lim-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
