package main

//sisi
import "fmt"

func GetGrade(a, b, c int) string {
	var middle int
	var grade string
	middle = (a + b + c) / 3
	switch {
	case 90 <= middle && middle <= 100:
		grade = "A"
	case 80 <= middle && middle < 90:
		grade = "B"
	case 70 <= middle && middle <= 80:
		grade = "C"
	case 60 <= middle && middle <= 70:
		grade = "D"
	case 0 <= middle && middle < 60:
		grade = "F"
	}
	return grade
}

func main() {
	result := 1
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	for i := 1; i <= n; i++ {
		result *= i
		fmt.Println(result)
	}
	fmt.Println(result)
	fmt.Println("Test")
}
