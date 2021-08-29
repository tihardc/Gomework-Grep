package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func checkString(line string, keyWord string, checkFunc checkFunc, formatter formatter) (string, error) {
	res, err := checkFunc(keyWord, line)
	if err != nil {
		return "", fmt.Errorf("error while checking the line: %w: %s", err, line)
	}
	if res == nil {
		return "", nil
	}
	return formatter(res, line), nil
}

func main() {

	var criterion checkFunc
	var format formatter

	config := mustGetConfig()

	file, err := os.Open(config.filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		switch strings.ToLower(config.option) {
		case "n":
			criterion, format = containsCheck, colorFormat
		case "y":
			criterion, format = containsCaseless, colorFormat
		case "r":
			criterion, format = containsRegEx, colorFormat
		case "f":
			criterion, format = containsRegEx, multicolorFormat
		}
		detected, err := checkString(string(line), config.keyString, criterion, format)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(detected)

	}

}
