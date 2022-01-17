package main

import "fmt"

var gender bool
var (
	name string
	age  int
)

var hobby string = "唱歌"
var score = 100
var height, university = 175, "清华大学隔壁八百公里外"

func main() {
	//fmt.Printf("name: %s , age: %d \n", name, age)

	a := 100
	fmt.Println(a)

	x, _ := doubleReturn()
	fmt.Println(x)
}

func doubleReturn() (x, y int) {
	return 1, 2
}
