package main

import "fmt"

func main() {
	slice1 := make([]string, 5)
	fmt.Println(slice1)
	slice2 := make([]int,3,5)
	fmt.Println(slice2)
	slice4 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice4)
	slice5 := []int{10, 20, 30}
	fmt.Println(slice5)
	slice6 := []string{99: ""}
	fmt.Println(slice6)

	slice := []int{10, 20, 30, 40}
	slice[1] = 22
	newSlice := slice[1:3]
	newSlice[0] = 33
	newSlice = append(newSlice, 66)
	newSlice = append(slice, 50)
	fmt.Println(slice)
	fmt.Println(newSlice)

	s1 := []int{1,2}
	s2 := []int{3,4}
	fmt.Printf("%v\n", append(s1, s2...))

	for index, value := range slice {
		fmt.Printf("인덱스 : %d   값 %d\n", index, value)
		fmt.Printf("값: %d  값의 주소: %X   원소의 주소%X\n", value, &value, &slice[index])
	}

	for index :=2; index < len(slice); index++ {
		fmt.Printf("인덱스 : %d   값 %d\n", index, slice[index])
	}

	slice7 := [][]int{{10}, {100, 200}}
	slice7[0] = append(slice7[0], 20)
	fmt.Println(slice7)




}
