package main

import (
	"strings"
	"regexp"
)

type (
	checkFunc func(keyWord, line string) ([]string, error)
)

func containsCheck(keyWord, line string) ([]string, error) {
	// TODO: add test
	if strings.Contains(line, keyWord) {
		return []string{keyWord}, nil
	}
	return nil , nil
}

func containsCaseless(keyWord, line string) ([]string, error) {
	re := regexp.MustCompile(`(?i)`+keyWord)
	return re.FindAllString(line, -1), nil
	}

func containsRegEx(keyWord, line string) ([]string, error) {
	re := regexp.MustCompile(keyWord)
	return re.FindAllString(line, -1), nil
	}