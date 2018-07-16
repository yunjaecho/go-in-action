package main

import "fmt"

func main() {
	colors := map[string]string {
		"AliceBlue" : "#f08234ff",
		"Coral" : "#f08324ff",
		"DarkGray" : "#f0432ff",
		"ForestGeen" : "#f4654ff",
	}

	for key, value := range colors {
		fmt.Printf("키 : %s , 값 : %s\n", key, value)
	}

	removeColor(colors, "Coral")

	fmt.Println("=========================")

	for key, value := range colors {
		fmt.Printf("키 : %s , 값 : %s\n", key, value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
