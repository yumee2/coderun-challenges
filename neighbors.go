package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func neighbors() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	strValues := strings.Fields(line)
	intsValues := make([]int, 0, len(strValues))
	for _, s := range strValues {
		i, _ := strconv.Atoi(s)
		intsValues = append(intsValues, i)
	}

	var counter int

	for i := 1; i < len(intsValues)-1; i++ {
		if intsValues[i] > intsValues[i-1] && intsValues[i] > intsValues[i+1] {
			counter++
		}
	}
	writer.WriteString(strconv.Itoa(counter))
	writer.WriteByte('\n')
}
