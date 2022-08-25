package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a==1")
		// break -> Go에서는 break를 생략해도 동일하게 동작한다.
	case 2:
		fmt.Println("a==2")
	case 3:
		fmt.Println("a==3")
		fallthrough // break 되지 않고 다음 case 까지 진행하고 싶을 경우 사용.
	case 4:
		fmt.Println("a==4")
	default:
		fmt.Println("a != 1, 2, 3")
	}
}
