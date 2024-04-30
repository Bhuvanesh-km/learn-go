package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world!")
	//declaring a variable
	var x int = 5
	fmt.Println(x)
	//declaring a variable without specifying the type
	y := 10
	fmt.Println(y)
	//type is inferred from the value
	var z = "Hello"
	fmt.Println(z)
	//declaration without initialization
	var i string
	var j int
	var k bool
	fmt.Println("i", i, "j", j, "k", k)
	//multiple declarations
	var l, m, n int = 1, 2, 3
	fmt.Println(l, m, n)
	var o, p = 1, "Hello"
	fmt.Println(o, p)
	//short hand declaration
	q, r := 1, 2
	fmt.Println(q, r)

	//declaring multiple variables
	var (
		a = 1
		b = 2
		c = 3
	)
	fmt.Println(a, b, c)
	//declaring constants
	const PI = 3.14
	fmt.Println(PI)
	//array declaration
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	//array declaration without specifying the size
	var arr2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	//slices
	var slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	//slices with make
	var slice2 = make([]int, 5)
	fmt.Println(slice2)

}
