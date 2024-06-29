package main

import "fmt"

func main() {
	// var ravjot = [3]int{1, 2, 3}
	// fmt.Println(ravjot)

	var a []int
	var b = []int{1, 2}
	fmt.Println(a)
	fmt.Println(b)

	a = append(a, 1)
	b = append(b, 1)
	fmt.Println(a)
	fmt.Println(b)

	a = b // it is overwiriting a baby

	fmt.Println(a)
	fmt.Println(b)

}
