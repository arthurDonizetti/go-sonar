package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a + b
}

func sumx(a int, b int) int {
	return a + b + a
}

func sumy(a int, b int) int {
	return a + b + b
}

func timesB(b int) int {
	return b * b
}

func timesA(a int) int {
	return a * a
}
