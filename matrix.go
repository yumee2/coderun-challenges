//go:build problem9

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

	line, _ := reader.ReadString('\n')
	strValues := strings.Fields(line)

	n, m, k := convertToInt(strValues[0]), convertToInt(strValues[1]), convertToInt(strValues[2])

	matrix1 := make([][]int, n)
	matrix2 := make([][]int, m)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		strValues := strings.Fields(line)

		matrix1[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix1[i][j] = convertToInt(strValues[j])
		}
	}

	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		strValues := strings.Fields(line)

		matrix2[i] = make([]int, k)
		for j := 0; j < k; j++ {
			matrix2[i][j] = convertToInt(strValues[j])
		}
	}

	result := make([][]int, k)

	for i := 0; i < k; i++ {
		result[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for h := 0; h < k; h++ {
				result[h][i] = result[h][i] + matrix1[i][j]*matrix2[j][h] // multiplication

			}
		}
	}

	for i := 0; i < k; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}
