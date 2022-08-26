package main

import "fmt"

// const Y int = 3

func main() {
	// var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	// fmt.Println(t)

	// // 배열의 선언 예시
	// var nums [5]int
	// fmt.Println(nums)

	// days := [3]string{"monday", "tuesday", "wednesday"}
	// fmt.Println(days)

	// var temps [5]float64 = [5]float64{24.3, 26.7}
	// fmt.Println(temps)

	// var s = [5]int{1: 10, 3: 30}
	// fmt.Println(s)

	// x := [...]int{10, 20, 30}
	// fmt.Println(x)

	// const x int = 5
	// a := [x]int{1, 2, 3, 4, 5}

	// b := [Y]int{1, 2, 3}

	// var c [10]int

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// nums := [...]int{10, 20, 30, 40, 50} // [5]int{10, 20 ...}
	// nums[2] = 300
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		fmt.Println(i, v)
	}

}
