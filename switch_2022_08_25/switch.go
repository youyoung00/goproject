package main

import "fmt"

func getMyAge() int {
	return 35
}

func main() {

	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is ", age)
	}

	// fmt.Println("age is ", age)

	// a := 4

	// switch a {
	// case 1:
	// 	fmt.Println("a == 1")
	// case 2:
	// 	fmt.Println("a == 2")
	// case 3:
	// 	fmt.Println("a == 3")
	// default:
	// 	fmt.Println("a != 1, 2, 3", a)
	// }

	// day := "monday"

	// switch day {
	// case "monday", "tuesday":
	// 	fmt.Println("월, 화요일은 수업 가는 날입니다.")
	// case "wendsday", "thursday", "friday":
	// 	fmt.Println("수, 목, 금요일은 실습 가는 날입니다.")
	// }

	// temp := 26

	// switch true {
	// case temp < 10, temp > 30:
	// 	fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다. ")
	// case temp >= 10 && temp < 20:
	// 	fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요. ")
	// case temp >= 15 && temp < 25:
	// 	fmt.Println("야외 활동하기 좋은 날씨입니다. ")
	// default:
	// 	fmt.Println("따뜻합니다. ")
	// }

}
