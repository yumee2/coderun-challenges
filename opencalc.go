package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	numbers := map[string]struct{}{}

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	line, _ := reader.ReadString('\n')
	strValues := strings.Fields(line)

	numbers[strValues[0]] = struct{}{}
	numbers[strValues[1]] = struct{}{}
	numbers[strValues[2]] = struct{}{}

	line2, _ := reader.ReadString('\n')
	value := strings.TrimSpace(line2)

	var count int

	for _, ch := range value {
		if _, exists := numbers[string(ch)]; exists {
			continue
		} else {
			numbers[string(ch)] = struct{}{}
			count++
		}

	}

	fmt.Println(count)
}
