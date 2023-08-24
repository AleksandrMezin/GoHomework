package main

import "fmt"

func main() {
	var a int
	var b int
	var sum int
	fmt.Scan(&a)
	for a < 5 {
		fmt.Println(a)
		a++
	}

	fmt.Scan(&b)
	if 10 <= b && b%8 == 0 && b <= 99 {
		sum += b
	}
	fmt.Println(sum)
}
