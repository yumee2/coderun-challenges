package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func triangle() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	readInt := func() int {
		line, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(line))
		return num
	}

	a, b, c := readInt(), readInt(), readInt()

	switch {
	case a+b <= c:
		fmt.Println("NO")
	case a+c <= b:
		fmt.Println("NO")
	case b+c <= a:
		fmt.Println("NO")
	default:
		fmt.Println("YES")
	}
}
