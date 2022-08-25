package main

import "fmt"

// var cnt int = 0 // 패키지 전역 변수

// func IncreaseAndReture() int {
// 	fmt.Println("IncreaseAndReture() 실행됨", cnt)
// 	cnt++
// 	return cnt
// }

func HasRichFriend(friend bool) bool {
	if friend {
		fmt.Println("친구들 중 부자가 있습니다.")
		return true
	} else {
		fmt.Println("친구들 중 부자가 없습니다.")
		return false
	}
}

func GetFriendsCount(friends int) int {
	fmt.Println("친구 숫자는 ", friends, "명입니다.")
	return friends
}

func main() {
	fmt.Println("[Goleng의 if문 학습 시작!!]")
	fmt.Println()

	price := 25000

	if price >= 50000 {
		if HasRichFriend(false) {
			fmt.Println("앗 신발끈이 풀렸네")
		} else {
			fmt.Println("나누어 내자")
		}
	} else if price >= 30000 {
		if GetFriendsCount(4) > 3 {
			fmt.Println("어이쿠 신발끈이....")
		} else {
			fmt.Println("사람도 얼마 없는데 나누어 내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다!!")
	}

	// temp := 17

	// var funcToInt int = IncreaseAndReture()
	// var boolVal bool = true

	// if temp > 28 {
	// 	fmt.Println("에어컨을 켠다")
	// } else if temp <= 3 {
	// 	fmt.Println("히터를 켠다")
	// } else if temp <= 18 {
	// 	fmt.Println("나가자 ! ")
	// } else {
	// 	fmt.Println("덥다")
	// }

	// var age = 22

	// if age >= 10 && age <= 15 {
	// 	fmt.Println("You are young")
	// } else if age > 30 || age < 20 {
	// 	fmt.Println("You are not 20s")
	// } else {
	// 	fmt.Println("Best age")
	// }
	// if true || IncreaseAndReture() < 5 {
	// 	fmt.Println("1 증가")
	// }

	// 누가낼 것인가? 예제

	// if price == 51000 &&
}
