package helpers

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func CommaSeperatedDay(input string) []string {
	return strings.Split(input, ", ")
}

func Md5String(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
