package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)
func main(){
	f, err := os.Open("day3/data-input3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	regex, _ := regexp.Compile(`mod\([0-9]+,[0-9]+\)`)
	r := bufio.NewReader(f)
	for {
		line, errInput := r.ReadString('\n')
		matches := regex.FindAllString(line,-1)
		for _, match:= range matches{
			var numStr1 string
			var numStr2 string
			var int1 int
			var int2 int
			for i:=4 ; i<len(match)
		}	

		
		if errInput != nil {
			break // End of file or read error
		}
		
	}
}