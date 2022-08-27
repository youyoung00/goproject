package main

import (
	"fmt"
	"strings"
	// "unsafe"
)

func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	// var rst string
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {

	fmt.Println(ToUpper1("abc"))
	fmt.Println(ToUpper2("abcd"))

	// var str string = "Hello World"
	// var slice []byte = []byte(str)

	// slice[2] = 'a'

	// stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	// sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	// fmt.Println(s)
	// fmt.Printf("str: \t%x\n", stringHeader.Data)
	// fmt.Printf("slice: \t%x\n", sliceHeader.Data)

	// str1 := "Hello 월드"
	// str2 := str1

	// stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	// stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	// fmt.Println(stringHeader1)
	// fmt.Println(stringHeader2)

	// fmt.Println("문자열 시작 !!")
	// 	poet1 := "죽는 날까지 하늘을 우러러\n한점 부끄럼이 없기를\n잎새이는 바람에도\n나는 괴로워했다."
	// 	poet2 := `죽는 날까지 하늘을 우러러
	// 한점 부끄럼이 없기를
	// 잎새이는 바람에도
	// 나는 괴로워했다.`

	// 	fmt.Println(poet1)
	// 	fmt.Println(poet2)

	// rune -> int32 별칭타입

	// arr := []rune(str) // 슬라이스 (동적 배열)

	// for _, v := range str {
	// 	fmt.Printf("타입: %T 값: %d 문자값: %c\n", v, v, v)
	// }

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("타입: %t 값: %d 문자값: %c\n", arr[i], arr[i], arr[i])
	// }

	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("타입: %t 값: %d 문자값: %c\n", str[i], str[i], str[i])
	// }
}
