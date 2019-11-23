package day2

import (
	"advent/utils"
	"fmt"
	"strings"
)

func Part2() {
	s := utils.GetStringSlice("./day2/input.txt")
	sMatrix := stringMatrix(s)

	var theTwo = [2]string{}
	for _, v := range sMatrix {
		theTwo = checkForTheOne(sMatrix, v)
		if theTwo[0] != "" {
			break
		}
	}
	fmt.Println(theTwo)
}

func checkForTheOne(sMatrix [][]string, b []string) [2]string {
	var theTwo = [2]string{}
	for _, a := range sMatrix {
		count := countDiffs(a, b)
		if count == 1 {
			theTwo[0] = strings.Join(a, "")
			theTwo[1] = strings.Join(b, "")
			break
		}
	}

	return theTwo
}

func countDiffs(a, b []string) int {
	var count = 0
	for i, _ := range a {
		if a[i] != b[i] {
			count = count + 1
		}
		if count > 1 {
			break
		}
	}

	return count
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
