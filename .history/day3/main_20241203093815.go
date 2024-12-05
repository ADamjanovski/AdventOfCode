package main

import (
	"fmt"
	"os"
	"regexp"
)
func main(){
	f, err := os.Open("day2/data-input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	r := bufio.NewReader(f)
}