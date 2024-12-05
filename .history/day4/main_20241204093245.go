package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	wholeSum := 0
	file, err := os.Open("day4/data-input4.txt")
	if err != nil {
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
	}
	stringData := string(data)
	str
}