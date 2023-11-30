package main

import (
	"crypto/md5"
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/util"
)

func run() error {
	input, err := util.ReadFileString("input")
	if err != nil {
		return err
	}

	for i := 0; ; i++ {
		sum := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		if sum[0] == 0 && sum[1] == 0 && sum[2] < 16 {
			fmt.Println("First 5 zeros:", i)
			break
		}
	}

	for i := 0; ; i++ {
		sum := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		if sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
			fmt.Println("First 6 zeros:", i)
			break
		}
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
