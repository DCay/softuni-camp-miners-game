package io

import (
	"bufio"
)

func ReadLine(consoleReader bufio.Reader) string {
	inputLine, err := consoleReader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return inputLine
}