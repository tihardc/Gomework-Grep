package main

import (
	"fmt"
	"strings"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	yellow = "\033[33m"
	green  = "\033[32m"
	blue   = "\033[34m"
)

type (
	formatter func(res []string, line string) string
)

// This assigns an integer value to a string in such a way that strings that are
// the same modulo case get the same value. It's a bit of a kludge but meh.

func foldedIndex(s string) int {
	s = strings.ToLower(s)
	i := 0
	for _, r := range s {
		i = i + int(r)
	}
	return i
}

func multicolorFormat(hits []string, line string) string {

	rainbow := []string{red, yellow, green, blue}

	for _, hit := range hits {
		hilight := rainbow[foldedIndex(hit)%len(rainbow)]
		line = strings.ReplaceAll(line, hit, fmt.Sprintf("%s%s%s", hilight, hit, reset))
	}
	return fmt.Sprintf("%s\n", line)
}

func colorFormat(hits []string, line string) string {
	for _, hit := range hits {
		line = strings.ReplaceAll(line, hit, fmt.Sprintf("%s%s%s", red, hit, reset))
	}
	return fmt.Sprintf("%s\n", line)
}
