package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)
func main(){
	f, err := os.Open("day3/data-input3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	regex1,_ := regexp.Compile(`do\(\)(.*?)don't\(\)`)
	regex2, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	r := bufio.NewReader(f)
	//PART ONE
	// sum:=0
	// for {
	// 	line, errInput := r.ReadString('\n')
	// 	matches := regex2.FindAllString(line,-1)
	// 	for _, match:= range matches{
	// 		var numStr string
	// 		var int1 int
	// 		var int2 int
	// 		fmt.Println(match)

	// 		for i:=4 ; i<len(match) ; i++{
	// 			if(match[i]==','){
	// 				int1, err= strconv.Atoi(numStr)
	// 				numStr=""
	// 				continue
	// 			}
	// 			if(match[i]==')'){
	// 				int2,err = strconv.Atoi(numStr)
	// 				break
	// 			}
	// 			numStr= numStr+ string(match[i])
	// 		}
	// 		sum+=(int1*int2)
	// 	}	

		
	// 	if errInput != nil {
	// 		break // End of file or read error
	// 	}
		
	// }
	// fmt.Println(sum)
	//PART TWO
	partTwoSum:=0
	for {
		line, errInput := r.ReadString('\n')
		result := regex1.FindAllString(line,-1)
		fmt.pr
		for _, res := range result{
			matches := regex2.FindAllString(res,-1)
			for _, match:= range matches{
				var numStr string
				var int1 int
				var int2 int
				
				for i:=4 ; i<len(match) ; i++{
					if(match[i]==','){
						int1, err= strconv.Atoi(numStr)
						numStr=""
						continue
					}
					if(match[i]==')'){
						int2,err = strconv.Atoi(numStr)
						break
					}
					numStr= numStr+ string(match[i])
				}
				partTwoSum+=(int1*int2)
			}	
		}
		lastOccurence := strings.LastIndex(line,"do()")


		if(!strings.Contains(line[lastOccurence:],"don't()")){
			matches := regex2.FindAllString(line[lastOccurence:],-1)
			for _, match:= range matches{
				var numStr string
				var int1 int
				var int2 int
				
				for i:=4 ; i<len(match) ; i++{
					if(match[i]==','){
						int1, err= strconv.Atoi(numStr)
						numStr=""
						continue
					}
					if(match[i]==')'){
						int2,err = strconv.Atoi(numStr)
						break
					}
					numStr= numStr+ string(match[i])
				}
				partTwoSum+=(int1*int2)
			}	
		}

		if errInput != nil {
			break // End of file or read error
		}
	}
	fmt.Println(partTwoSum)
}