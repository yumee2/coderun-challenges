package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tags() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	if n < 3 {
		fmt.Println(n)
		return
	}

	last, prelast := 1, 1

	counter := 2

	i := 3
	for i < n {
		val := last + prelast
		last, prelast = val, last
		i++

		counter += val
	}
	fmt.Println(counter + (last + prelast))
}
