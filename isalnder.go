package main

import (
	"fmt"
)

func main() {
	y0 := []int{1, 0, 0, 0, 0, 0}
	y1 := []int{0, 1, 0, 1, 1, 1}
	y2 := []int{0, 0, 1, 0, 1, 0}
	y3 := []int{1, 1, 0, 0, 1, 0}
	y4 := []int{1, 0, 1, 1, 0, 0}
	y5 := []int{1, 0, 0, 0, 0, 1}
	is := [][]int{y0, y1, y2, y3, y4, y5}

	fmt.Println(is[0][0])
	fmt.Println(is[1][1])
	fmt.Println(is[0][1])
	fmt.Println(is[4][1]) // (y,x)
	fmt.Println(len(is[1]))

	bordersafe := border_check(is)
	fmt.Println("Coordenadas Bordersafe")
	fmt.Println(bordersafe)
}

func border_check(is [][]int) []int {
	var bordersafe []int
	// check for top and bottom
	for i := 0; i < len(is); i++ {
		if is[0][i] == 1 {
			bordersafe = append(bordersafe, 0)
			bordersafe = append(bordersafe, i)

		}

		if is[len(is)-1][i] == 1 {
			bordersafe = append(bordersafe, len(is)-1)
			bordersafe = append(bordersafe, i)
		}
	}
	// check left and right

	for i := 1; i < len(is[1])-1; i++ {
		if is[i][0] == 1 {
			bordersafe = append(bordersafe, i)
			bordersafe = append(bordersafe, 0)

		}

		if is[i][len(is)-1] == 1 {
			bordersafe = append(bordersafe, i)
			bordersafe = append(bordersafe, len(is)-1)
		}
	}
	return bordersafe
}
