package main

import (
	"fmt"

	"github.com/darshanmadesh/GoPackagePractice/basicMath"
	"github.com/darshanmadesh/GoPackagePractice/basicMath/square"
	"github.com/darshanmadesh/GoPackagePractice/helloWorld"
)

func main() {
	helloWorld.PrintHelloWorld()

	fmt.Println("add function : ", basicMath.AddNums(25, 35) )

	fmt.Println("square function : ", square.SquareNum(12) )

}