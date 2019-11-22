package day1

import (
	"fmt"
	"strconv"
	"strings"

	"advent/utils"
)

func Part2() {
	numberSlice := getNumberSlice()

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
	numberSlice := getNumberSlice()
	var sum = 0
	for _, v := range numberSlice {
		sum = sum + v
	}

	fmt.Println(sum)
}

func getNumberSlice() []int {
	s, _ := utils.GetInput("./day1/input.txt")
	stringSlice := strings.Split(s, "\n")

	var numberSlice []int
	for _, v := range stringSlice {
		n, _ := strconv.Atoi(v)
		numberSlice = append(numberSlice, n)
	}

	return numberSlice[:len(numberSlice)-1]
}
