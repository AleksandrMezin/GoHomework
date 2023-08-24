package main
import "fmt"

package calc
type calc struct {
	
}


func main() {
	var i, j int
	c := 0
	for i = 1; i < 10; i++ {
		for j = 9; j > 0; j-- {
			switch {
			default:
				break
			case i == j:
				c := i + j
				c = c * 2
				continue
			}
			println(i, j)
		}
	}
	fmt.Println(c)
}
