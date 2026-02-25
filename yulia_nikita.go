package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func yulia_nikita() {

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	strValues := strings.Fields(line)

	x, _ := strconv.Atoi(strValues[0])
	y, _ := strconv.Atoi(strValues[1])

	result := x + y

	writer.WriteString(strconv.Itoa(result))
	writer.WriteByte('\n')
}
