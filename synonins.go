//go:build problem11

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
	line, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(line))

	dict := make(map[string]string)

	for i := 0; i < num; i++ {
		str, _ := reader.ReadString('\n')
		words := strings.Fields(strings.TrimSpace(str))

		dict[words[0]] = words[1]
		dict[words[1]] = words[0]
	}

	wrd, _ := reader.ReadString('\n')
	fmt.Println(dict[strings.TrimSpace(wrd)])
}
