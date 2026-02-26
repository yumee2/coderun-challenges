package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func counting_words() {
	words := map[string]struct{}{}

	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	data, _ := io.ReadAll(reader)
	text := string(data)
	strValues := strings.Fields(text)

	for _, value := range strValues {
		words[value] = struct{}{}
	}

	fmt.Println(len(words))
}
