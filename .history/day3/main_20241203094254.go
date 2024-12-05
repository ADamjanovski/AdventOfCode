package main

import (
	"bufio"
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
	regex, _ := regexp.Compile("mod([0-9]+,[0-9]+)")
	r := bufio.NewReader(f)
	for {
		line, errInput := r.ReadString('\n')
		matches,err := regex.FindAllString(line,-1)


		
		if errInput != nil {
			break // End of file or read error
		}
		
	}
}