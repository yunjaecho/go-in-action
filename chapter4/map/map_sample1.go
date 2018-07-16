package main

import "fmt"

func main() {
	colors := map[string]string{}
	colors["Red"] = "#da1337"

	var colors2 = map[string]string{}
	colors2["Red"] = "#da1337"

	value, exists := colors["Blue"]
	if exists {
		fmt.Println(value)
	}

	value2 := colors["Blue"]
	if value2 != "" {
		fmt.Println(value2)
	}


	colors3 := map[string]string {
		"AliceBlue" : "#f08234ff",
		"Coral" : "#f08324ff",
		"DarkGray" : "#f0432ff",
		"ForestGeen" : "#f4654ff",
	}

	for key, value := range colors3 {
		fmt.Printf("키 : %s , 값 : %s\n", key, value)
	}

	delete(colors3, "Coral")

	for key, value := range colors3 {
		fmt.Printf("키 : %s , 값 : %s\n", key, value)
	}
}
