package main

import (
	"fmt"
)

type Page struct {
	Title string
	Body string
}

func main() {
	var p1 = &Page{"darshan", "madesh"}
	fmt.Print(p1)
}