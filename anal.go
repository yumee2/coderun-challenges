package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func analytic() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	convertToInt := func(val string) int {
		num, _ := strconv.Atoi(strings.TrimSpace(val))
		return num
	}

	line, _ := reader.ReadString('\n')
	strValues := strings.Fields(line)

	a, b, c := convertToInt(strValues[0]), convertToInt(strValues[1]), convertToInt(strValues[2])

	dis := (b * b) - 4*a*c

	if dis < 0 {
		fmt.Println(0)
	} else if dis == 0 {
		x1 := float64(-b) / float64(2*a)
		fmt.Println(1)
		fmt.Printf("%.6f\n", x1)
	} else {
		x1 := (float64(-b) + math.Sqrt(float64(dis))) / float64(2*a)
		x2 := (float64(-b) - math.Sqrt(float64(dis))) / float64(2*a)
		fmt.Println(2)
		if x1 < x2 {
			fmt.Printf("%.6f %.6f\n", x1, x2)
		} else {
			fmt.Printf("%.6f %.6f\n", x2, x1)
		}
	}
}
