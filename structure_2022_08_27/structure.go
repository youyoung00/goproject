package main

import (
	"fmt"
	"unsafe"
)

// 구조체 정의
type User struct {
	A int8
	B int8
	C int8
	D int
	E int
	// Age   int32   // 4바이트
	// Score float64 // 8바이트
}

// type VIPUser struct {
// 	User
// 	VIPLevel int
// 	Price    int
// }

func main() {

	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))

	// var a int = 45
	// user := User{23, 77.2}
	// fmt.Println(unsafe.Sizeof(user), unsafe.Sizeof(a))

	// user := User{"송하나", "hana", 23}
	// vip := VIPUser{
	// 	User{"화랑", "hwarang", 48},
	// 	3,
	// 	250,
	// }

	// fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.Id, user.Age)
	// fmt.Printf("VIP유저: %s ID: %s 나이: %d\n VIP레벨: %d VIP가격: %d만원\n",
	// 	vip.Name,
	// 	vip.Id,
	// 	vip.Age,
	// 	vip.VIPLevel,
	// 	vip.Price,
	// )

	// var house House
	// house.Address = "서울시 강남구 ..."
	// house.Size = 28
	// house.Price = 10
	// house.Category = "아파트"

	// // 구조체의 활용
	// fmt.Printf("주소:%s 사이즈:%d평 가격:%v억원 종류:%s\n",
	// 	house.Address, house.Size, house.Price, house.Category)

	// // 초기화
	// var home House
	// fmt.Println(home)

	// var home2 House = House{"서울시 강동구", 28, 9.80, "아파트"}
	// fmt.Println(home2)

	// var home3 House = House{
	// 	"서울시 강동구",
	// 	28,
	// 	9.80,
	// 	"아파트",
	// }
	// fmt.Println(home3)

	// // 일부 요소만 초기화
	// var home4 House = House{Size: 28, Category: "아파트"}
	// fmt.Println(home4)

}
