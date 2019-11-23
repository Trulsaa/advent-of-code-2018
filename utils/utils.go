package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetInput(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	return string(b), nil
}

func GetStringSlice(path string) []string {
	s, _ := GetInput(path)
	stringSlice := strings.Split(s, "\n")
	return stringSlice[:len(stringSlice)-1]
}

func GetNumberSlice(path string) []int {
	stringSlice := GetStringSlice(path)

	var numberSlice []int
	for _, v := range stringSlice {
		n, _ := strconv.Atoi(v)
		numberSlice = append(numberSlice, n)
	}

	return numberSlice
}

type Timer struct {
	startTime time.Time
}

func (t *Timer) Start() {
	t.startTime = time.Now()
}

func (t *Timer) Stop() {
	stopTime := time.Now()
	elapsed := stopTime.Sub(t.startTime)
	fmt.Printf("Sorry for the %v wait\n", elapsed)
}
