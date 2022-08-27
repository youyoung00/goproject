package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg *Data) {
	arg.value = 999
	fmt.Printf("ChangeDataValue = %d\n", arg.value)

	arg.data[100] = 999
	fmt.Printf("ChangeDataData[100] = %d\n", arg.data[100])
}

func main() {

	var data Data

	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	// fmt.Println("포인터 시작 !!")
	// var a int = 500
	// var p *int

	// p = &a

	// fmt.Printf("p의 값: %p\n", p)
	// fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	// *p = 100
	// fmt.Printf("a의 값: %d\n", a)

	// var a int = 10
	// var b int = 20

	// var p1 *int = &a
	// var p2 *int = &a
	// var p3 *int = &b

	// fmt.Printf("p1 == p2: %v\n", p1 == p2)
	// fmt.Printf("p2 == p3: %v\n", p2 == p3)

	// var p *int
	// if p != nil {

	// }

}
