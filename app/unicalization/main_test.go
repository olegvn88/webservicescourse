package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

var testOk = `1
2
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))

	out := new(bytes.Buffer)
	err := unic(in, out)
	if err != nil {
		t.Errorf("test for Ok failed")
	}
	result := out.String()
	fmt.Print(result)
	if result != testOkResult {
		t.Errorf("test for Ok is failed - results is not matched\n %v %v", result, testOkResult)
	}
}

var testFail = `1
2
1`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))

	out := new(bytes.Buffer)
	err := unic(in, out)
	if err == nil {
		t.Errorf("test for Ok failed - error: %v", err)
	}
}

func TestLength(t *testing.T) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz123456789"
	b := make([]byte, 63)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	fmt.Println(string(b))
	fmt.Println(len(string(b)))
	fmt.Println(len("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa234aaaaabbbbcccccccc"))
}
