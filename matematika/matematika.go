package main

import "fmt"

func main() {

	var a = 50
	var b = 50
	var c = 30
	var d = 10
	var e = a + b*d/c

	fmt.Println(e)

	var i = 10

	i += 10 //i = i + 10
	fmt.Println(i)

	i += 5 //i = i + 5
	fmt.Println(i)

	var j = 1

	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)

	j-- // j = j - 1
	fmt.Println(j)

}
