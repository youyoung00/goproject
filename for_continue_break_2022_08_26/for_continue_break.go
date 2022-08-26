package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main() {

	a := 1
	b := 1

OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

	// stdin := bufio.NewReader((os.Stdin))
	// for {
	// 	fmt.Println("입력하세요")
	// 	var number int
	// 	_, err := fmt.Scanln(&number)
	// 	if err != nil {
	// 		fmt.Println("숫자로 입력하세요")

	// 		// 키보드 버퍼를 지웁니다.
	// 		stdin.ReadString('\n')
	// 		continue
	// 	}
	// 	fmt.Printf("입력하시 숫자는 %d입니다.\n", number)
	// 	if number%2 == 0 {
	// 		break // 짝수검사를 합니다.
	// 	}
	// }
	// fmt.Println("for문이 종료되었습니다.")

}
