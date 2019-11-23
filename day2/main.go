package day2

import (
	"advent/utils"
	"fmt"
	"strings"
)

func Part2() {

}

func Part1() {
	s := utils.GetStringSlice("./day2/input.txt")
	sMatrix := stringMatrix(s)
	fmt.Println(checksum(sMatrix))
}

func checksum(charsSlice [][]string) int {
	var double, tripple int
	for _, chars := range charsSlice {
		c := checkFoDoublesAndTripples(chars)
		double = double + c.double
		tripple = tripple + c.tripple
	}
	return double * tripple
}

type doublesAndTripples struct {
	double  int
	tripple int
}

func checkFoDoublesAndTripples(chars []string) doublesAndTripples {
	counts := map[string]int{}
	for _, char := range chars {
		counts[char] = counts[char] + 1
	}

	dsAndTs := doublesAndTripples{}
	for _, count := range counts {
		if count == 2 {
			dsAndTs.double = 1
		}
		if count == 3 {
			dsAndTs.tripple = 1
		}
	}

	return dsAndTs
}

func stringMatrix(slice []string) [][]string {
	var stringSlices [][]string
	for _, v := range slice {
		slice := strings.Split(v, "")
		stringSlices = append(stringSlices, slice)
	}
	return stringSlices
}
