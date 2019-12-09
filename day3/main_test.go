package day3

import (
	"strconv"
	"strings"
	"testing"
)

func TestDrawClaim(t *testing.T) {
	// fabric [1000][1000]int, claim claim
	// #1 @ 1,3: 4x4
	// #2 @ 3,1: 4x4
	// #3 @ 5,5: 2x2
	claims := []claim{
		claim{1, 1, 3, 4, 4},
		claim{2, 3, 1, 4, 4},
		claim{3, 5, 5, 2, 2},
	}

	expected := `........
...2222.
...2222.
.11XX22.
.11XX22.
.111133.
.111133.
........
`

	r := printEightByEight(drawFabric(claims))
	if r != expected {
		t.Errorf("Expected length: %v", len(expected))
		t.Errorf("Result length: %v", len(r))
		t.Errorf("Expected:\n%s", expected)
		t.Errorf("Result:\n%s", r)
	}
}

func printEightByEight(fabric [1000][1000]int) string {
	var s string
	for i, line := range fabric {
		s = s + strings.Join(intsToStrings(line[:8]), "") + "\n"
		if i >= 7 {
			break
		}
	}
	return s
}
