package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(time.Second) // 1초 동안 멈추어라
		fmt.Println(i)
		i++
	}
}
