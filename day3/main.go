package day3

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part2() {
	s := utils.GetStringSlice("./day3/input.txt")
	claims := parseClaims(s)
	fabric := drawFabric(claims)
	id := checkForOverlap(&fabric, claims)
	fmt.Println(id)
}

func Part1() {
	s := utils.GetStringSlice("./day3/input.txt")
	claims := parseClaims(s)
	bottomRightCoordinates := getBottomRightCooridnate(claims)
	fmt.Printf("Bottom right most coordinate: %v\n", bottomRightCoordinates)
	fabric := drawFabric(claims)
	count := countOverlaps(&fabric)
	fmt.Println(count)
}

type claim struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

func countOverlaps(fabric *[1000][1000]int) int {
	var count = 0
	for _, line := range fabric {
		for _, p := range line {
			if p < 0 {
				count++
			}
		}
	}
	return count
}

func checkForOverlap(fabric *[1000][1000]int, claims []claim) int {
	for _, claim := range claims {
		if isNotOverlapped(fabric, claim) {
			return claim.id
		}
	}
	return 0
}

func isNotOverlapped(fabric *[1000][1000]int, claim claim) bool {
	var isNoOverlap = true
	for y := claim.left; y <= claim.left+claim.width-1; y++ {
		for x := claim.top; x <= claim.top+claim.height-1; x++ {
			if fabric[x][y] != claim.id {
				isNoOverlap = false
			}
		}
	}

	return isNoOverlap
}

func drawFabric(claims []claim) [1000][1000]int {
	var fabric = &[1000][1000]int{}
	for _, claim := range claims {
		drawClaim(fabric, claim)
	}
	return *fabric
}

func forFabric(fabric *[1000][1000]int, claim claim, f func(coordinate int, id int)) {
	for y := claim.left; y <= claim.left+claim.width-1; y++ {
		for x := claim.top; x <= claim.top+claim.height-1; x++ {
			f(fabric[x][y], claim.id)
		}
	}
}

func drawClaim(fabric *[1000][1000]int, claim claim) {
	for y := claim.left; y <= claim.left+claim.width-1; y++ {
		for x := claim.top; x <= claim.top+claim.height-1; x++ {
			if fabric[x][y] == 0 {
				fabric[x][y] = claim.id
			} else {
				fabric[x][y] = -1
			}
		}
	}
}

func getBottomRightCooridnate(claims []claim) [2]int {
	coordinates := [2]int{}
	for _, claim := range claims {
		x := claim.left + claim.width
		if coordinates[0] < x {
			coordinates[0] = x
		}
		y := claim.top + claim.height
		if coordinates[1] < y {
			coordinates[1] = y
		}
	}
	return coordinates
}

// #32 @ 863,904: 22x20
func parseClaims(s []string) []claim {
	var claims []claim
	for _, v := range s {
		claims = append(claims, parseClaim(v))
	}
	return claims
}

func parseClaim(s string) claim {
	slice := strings.Split(s, " ")

	values := [5]string{}
	values[0] = strings.Replace(slice[0], "#", "", 1)
	values[1] = strings.Split(slice[2], ",")[0]
	values[2] = strings.Replace(strings.Split(slice[2], ",")[1], ":", "", 1)
	values[3] = strings.Split(slice[3], "x")[0]
	values[4] = strings.Split(slice[3], "x")[1]

	var numbers []int
	for _, v := range values {
		n, _ := strconv.Atoi(v)
		numbers = append(numbers, n)
	}

	return claim{
		numbers[0],
		numbers[1],
		numbers[2],
		numbers[3],
		numbers[4],
	}
}

func printFabric(fabric [1000][1000]int) string {
	var s string
	for _, line := range fabric {
		s = s + strings.Join(intsToStrings(line[:1000]), "") + "\n"
	}
	return s
}

func intsToStrings(ints []int) []string {
	ss := []string{}
	for _, v := range ints {
		switch v {
		case 0:
			ss = append(ss, ".")
		case -1:
			ss = append(ss, "X")
		default:
			ss = append(ss, strconv.Itoa(v))
		}
	}
	return ss
}
