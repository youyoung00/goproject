package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func main() {
	fmt.Println("My favorite color is", colorToString(Blue))
	fmt.Println("My favorite color is", colorToString(2))
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}
