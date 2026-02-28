//go:build problem12

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	var slice []int

	addValToSlice := func(val string) int {
		num, _ := strconv.Atoi(strings.TrimSpace(val))
		slice = append(slice, num)
		return num
	}

	reader.ReadString('\n')

	line2, _ := reader.ReadString('\n')
	strValues := strings.Fields(line2)
	for _, val := range strValues {
		addValToSlice(val)
	}

	line3, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(line3))

	var result int
	dif := math.MaxInt
	for _, val := range slice {
		if math.Abs(float64(val-num)) < float64(dif) {
			dif = int(math.Abs(float64(val - num)))
			result = val
		}
	}

	fmt.Println(result)
}
