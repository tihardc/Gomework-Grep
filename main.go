package main

import (
	"log"
	"io"
	"bufio"
	"fmt"
	"os"
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
		switch config.option {
		case "n", "N": criterion = containsCheck
		case "y", "Y": criterion = containsCaseless
		case "r", "R": criterion = containsRegEx
		}
		detected, err := checkString(string(line), config.keyString, criterion, colorFormat)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(detected)

	}


}