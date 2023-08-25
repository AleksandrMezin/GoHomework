package main

import "fmt"

func read(message string) string {
	return message
}

func plus(a, b int) {
	fmt.Println(a + b)
}

func main() {
	fmt.Println(read("Dudu and Pupu"))
	plus(13, 54)
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)
	fmt.Printf("%T", x)

	fmt.Println("go go go")
}
