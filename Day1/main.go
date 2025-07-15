package main

import (
	"fmt"
	"os"
	"strings"
)

func KtoC(value float32) float32 {
	return value - 273.15
}
func CtoK(value float32) float32 {
	return value + 273.15
}

func main() {
	var name string
	var util string
	var value float32
	var celsius = "celsius"
	var kelvin = "kelvin"

	fmt.Println("Your name: ")
	fmt.Fscan(os.Stdin, &name)

	fmt.Println("Temperature in Celsius or in Kelvin?: ")
	fmt.Fscan(os.Stdin, &util)

	str := strings.TrimSpace(strings.ToLower(util))

	if str != kelvin && str != celsius {
		fmt.Println("Please choose a correct unit (celsius or kelvin).")
		return
	}

	fmt.Println("Temperature value: ")
	fmt.Fscan(os.Stdin, &value)

	if str == celsius {
		fmt.Print("Hi, ", name, "Celsius ", value, " in Kelvin = ")
		fmt.Printf("%.3f\n", CtoK(value))
	} else if str == kelvin {
		fmt.Print("Hi, ", name, "Kelvin ", value, " in Celsius = ")
		fmt.Printf("%.3f\n", KtoC(value))
	}
}
