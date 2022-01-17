package main

const PI = 3.1415926
const E = 2.71828
const (
	pi = 3.1415926
	e  = 2.71828
	a
	b
)
const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)
const (
	m1 = iota //0
	m2        //1
	_
	m4 //3
)

const (
	k1 = iota //0
	k2 = 100  //100
	k3 = iota //2
	k4        //3
)
const k5 = iota //0

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	l1, l2 = iota + 1, iota + 2 //1,2
	l3, l4                      //2,3
	l5, l6                      //3,4
)

func main() {
}
