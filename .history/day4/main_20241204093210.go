package main

import (
	"fmt"
	"os"
)

func main(){
	wholeSum := 0
	file, err := os.Open("day3/data-input3.txt")
	if err != nil {
		return
	}
	defer file.Close()
	
}