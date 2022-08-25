package main

import "fmt"

const Pig int = 0
const Cow int = 1
const Chicken int = 2

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

func iotaTester() {
	const (
		C1 int = iota + 1
		C2
		C3
		C4
		C5
	)

	const (
		A1 int = iota
		A2
		A3
		A4
		A5
	)

	fmt.Println(C5)
	fmt.Println(C4)
	fmt.Println(C3)
	fmt.Println(C2)
	fmt.Println(C1)

	// fmt.Println("\n")

	fmt.Println(A1)
	fmt.Println(A2)
	fmt.Println(A3)
	fmt.Println(A4)
	fmt.Println(A5)
}

func main() {
	PrintAnimal(Cow)
	PrintAnimal(Pig)
	PrintAnimal(7)

	iotaTester()
}
