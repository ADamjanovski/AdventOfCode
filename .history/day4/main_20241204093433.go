package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){
	wholeSum := 0
	file, err := os.Open("day4/data-input4.txt")
	if err != nil {
		return
	}
	defer file.Close()
	r := bufio.NewReader(f)

	stringData := string(data)
	var charsMatrix [][] string
	for{
		line, err := r.
	}
}