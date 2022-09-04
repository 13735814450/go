package main

import (
	"config"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	var j int = 0
	j = t1(10, 1)
	fmt.Println(j)
	config.LoadConfig()
}

func t1(i int, j int) int {
	var k int
	k = i + j
	return k
}
