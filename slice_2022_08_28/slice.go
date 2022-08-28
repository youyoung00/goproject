package main

import "fmt"

func main() {

	var sumSlice = []int{1, 2, 3}
	var multiplySlice = []int{1, 2, 3}

	for i := 0; i < len(sumSlice); i++ {
		sumSlice[i] += 10
	}

	fmt.Println(sumSlice)

	// for i := 0; i < len(multiplySlice); i++ {
	// 	multiplySlice[i] = i * 2
	// }
	for i, v := range multiplySlice {
		multiplySlice[i] = v * 2
	}
	fmt.Println(multiplySlice)

	// var makeSlice = make([]int, 3) // 초기화 방법1
	// makeSlice2 := []int{0, 0, 0}   // 초기화 방법2

	// if len(slice) == 0 {
	// 	fmt.Println("slice is empty", slice)
	// }

	// slice[1] = 10
	// fmt.Println(slice)

	// fmt.Println(makeSlice)
	// fmt.Println(makeSlice2)

}
