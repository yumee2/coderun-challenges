//go:build problem8

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
	var nums []int

	line, _ := reader.ReadString('\n')
	strValues := strings.Fields(line)

	for i, val := range strValues {
		num, _ := strconv.Atoi(strings.TrimSpace(val))
		if len(nums) > 0 {
			if num <= nums[i-1] {
				fmt.Println("NO")
				return
			}
		}
		nums = append(nums, num)
	}

	fmt.Println("YES")
}
