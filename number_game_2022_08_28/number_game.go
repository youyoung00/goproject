package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader((os.Stdin))

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)
	cnt := 1
	for {
		fmt.Println("숫자를 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else {
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췃습니다. 축하합니다. 시도한 횟수: ", cnt)
				break
			}
			cnt++
		}
	}
	// loc, _ := time.LoadLocation(("Asia/Seoul"))
	// const LongForm = "Jan 2, 2006 at 3:04pm"
	// t1, _ := time.ParseInLocation(LongForm, "Jun 14, 2021 at 10:00pm", loc)
	// fmt.Println(t1, t1.Location(), t1.UTC())
	// fmt.Println()

	// const shortForm = "2006-Jan-02"
	// t2, _ := time.Parse(shortForm, "2021-Jun-14")
	// fmt.Println(t2, t2.Location())
	// fmt.Println()

	// t3, err := time.Parse("2021-06-01 15:20:21", "2021-06-14 20:04:05")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(t3, t3.Location())
	// fmt.Println()

	// d := t2.Sub(t1)
	// fmt.Println(d)

	// t := time.Now()
	// rand.Seed(t.UnixNano())

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(rand.Intn(100))
	// }
}
