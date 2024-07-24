package helpers

import (
	"strings"
)

func CommaSeperatedDay(input string) []string {
	return strings.Split(input, ", ")
}
