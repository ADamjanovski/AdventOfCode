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
	rexe, _ := regexp.Compile("p([a-z]+)ch")
	r := bufio.NewReader(f)
	for {
		line, errInput := r.ReadString('\n')
		arr := regexp.FindAllString("")
	}
}