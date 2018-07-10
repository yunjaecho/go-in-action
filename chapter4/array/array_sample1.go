package main

import "fmt"

func main() {
	array := [5]int{10, 20, 30, 40, 50}
	//array := [...]int{10, 20, 30, 40, 50}
	// 각원소는 지정된 값으로초기화된다. 배열의 나머지 원소는 제로 값으로 초기화 된다.
	//array := [5]int{1,2,3}
	fmt.Println(array)
	fmt.Println(array[2])

	// 다섯 개의 원소를 가지는 정수 포인터 배열을 선언한다.
	// 인덱스 0과 1을 정수 포인터로 초기화 한다.
	array_point := [5]*int{0: new(int), 1:new(int)}
	*array_point[0] = 10
	*array_point[1] = 20

	var array1 [5]string
	array2 := [5]string{"red", "blue", "green", "yellow", "pink"}
	// array2 의 값을 array1로 복사
	array1 = array2

	fmt.Println(array1)
	fmt.Println(array2)
	array2[1] = "red2"
	fmt.Println(array1)
	fmt.Println(array2)

	fmt.Println("=====================================")
	fmt.Println("=========== Array Pointer ===========")
	fmt.Println("=====================================")

	var array3 [3]*string
	array4 := [3]*string{new(string), new(string), new(string)}
	*array4[0] = "Red"
	*array4[1] = "Blue"
	*array4[2] = "Green"

	// 동일한 값을 가지는 두개의 포인트 배열
	array3 = array4

	fmt.Println(*array3[1])
	fmt.Println(*array4[1])
	*array4[1] = "Black"
	fmt.Println(*array3[1])
	fmt.Println(*array4[1])


}
