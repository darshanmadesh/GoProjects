package main

import "fmt"

func main() {
	var slice = []int{3,2,4,1,6}

	a := slice[0:3]
	b := slice[3:5]

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 100

	fmt.Println(slice)


}