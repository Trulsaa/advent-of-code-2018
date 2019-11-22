package utils

import (
	"fmt"
	"io/ioutil"
	"os"
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
