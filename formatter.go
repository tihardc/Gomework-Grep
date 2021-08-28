package main

import (
	"fmt"
	"strings"
)

const (
	reset = "\033[0m"
	red   = "\033[31m"
)

type (
	formatter func(res []string, line string) string
)

func colorFormat(hits []string, line string) string {
	for _, hit := range hits {
		line = strings.ReplaceAll(line, hit, fmt.Sprintf("%s%s%s", red, hit, reset))
	}
	return fmt.Sprintf("%s\n", line)
}
