package main

import (
	"fmt"

	"github.com/harikaduyu/aoc-skeleton-go/utils"
)

func result(input string, part int) int {
	fmt.Println(input, part)
	fmt.Print("Your solution here:")

	return 0
}

func main() {
	input := utils.ReadInput()
	result_1 := result(input, 1)
	fmt.Println("Part1:", result_1)
	result_2 := result(input, 2)
	fmt.Println("Part2:", result_2)
}
