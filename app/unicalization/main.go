package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	err := unic(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}

func unic(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}

		if txt < prev {
			return fmt.Errorf("file is not sorted")
		}

		prev = txt
		fmt.Fprintln(output, txt)
	}
	return nil
}
