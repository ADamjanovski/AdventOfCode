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
	regex, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	r := bufio.NewReader(f)
	sum:=0
	for {
		line, errInput := r.ReadString('\n')
		matches := regex.FindAllString(line,-1)
		for _, match:= range matches{
			var numStr string
			var int1 int
			var int2 int
			fmt.Println(match)

			for i:=4 ; i<len(match) ; i++{
				if(match[i]==','){
					fmt.Println(numStr)
					int1, err= strconv.Atoi(numStr)
					fmt.Println(int1)
					numStr=""
				}
				if(match[i]==')'){
					int2,err = strconv.Atoi(numStr)
					break
				}
				numStr= numStr+ string(match[i])
			}
			sum+=(int1*int2)
		}	

		
		if errInput != nil {
			break // End of file or read error
		}
		
	}
	fmt.Println(sum)
}