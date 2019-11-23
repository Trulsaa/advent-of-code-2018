package day1

import (
	"fmt"

	"advent/utils"
)

func Part2() {
	numberSlice := utils.GetNumberSlice("./day1/input.txt")

	var sums = []int{numberSlice[0]}

	var i = 1
	for {
		sum := sums[len(sums)-1] + numberSlice[i%len(numberSlice)]
		if includes(sums, sum) {
			fmt.Println(sum)
			break
		}
		sums = append(sums, sum)
		i = i + 1
	}
}

func includes(slice []int, i int) bool {
	for _, v := range slice {
		if v == i {
			return true
		}
	}
	return false
}

func Part1() {
	numberSlice := utils.GetNumberSlice("./day1/input.txt")
	var sum = 0
	for _, v := range numberSlice {
		sum = sum + v
	}

	fmt.Println(sum)
}
