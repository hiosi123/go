package main

import "fmt"

func main() {
	var a int
	var b int
	// var f float64 = 321.123
	// fmt.Print("a:", a, "b:", b, "c:", f)
	// fmt.Printf("a: %d b: %d f:%f \n", a, b, f)

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
