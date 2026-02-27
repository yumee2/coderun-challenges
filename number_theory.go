//go:build problem10

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	convertToInt := func(val string) int {
		num, _ := strconv.Atoi(strings.TrimSpace(val))
		return num
	}

	findMaxMin := func(a, b int) (int, int) {
		if a > b {
			return a, b
		}
		return b, a
	}

	line, _ := reader.ReadString('\n')
	strValues := strings.Fields(line)

	a, b := convertToInt(strValues[0]), convertToInt(strValues[1])
	max, min := findMaxMin(a, b)

	var nod, nok int

	if max%min == 0 {
		nod = min
		nok = max
	} else {
		nod = gcd(max, min)

		nok = max * min / nod
	}

	fmt.Println(nod, nok)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
