package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hash(doorID string, num int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(doorID+strconv.Itoa(num))))
}

func password(doorID string) string {
	result := ""
	for i := 0; ; i++ {
		h := hash(doorID, i)
		if strings.HasPrefix(h, "00000") {
			result += h[5:6]
			if len(result) == 8 {
				return result
			}
		}
	}
}

func password2(doorID string) string {
	result := []byte("________")
	toFind := 8
	for i := 0; ; i++ {
		h := hash(doorID, i)
		if strings.HasPrefix(h, "00000") {
			pos := int(h[5] - '0')
			if pos > 7 {
				continue
			}
			if result[pos] != '_' {
				continue
			}

			result[pos] = h[6]
			toFind--
			if toFind == 0 {
				return string(result)
			}
		}
	}
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
