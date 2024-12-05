package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)
func findMulOfMatches(result string)int{
	sum:=0
	regex2, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	matches := regex2.FindAllString(result,-1)
	for _, match:= range matches{
		var numStr string
		var int1 int
		var int2 int
		var err error
		for i:=4 ; i<len(match) ; i++{
			if(match[i]==','){
				int1, err= strconv.Atoi(numStr)
				numStr=""
				if(err!=nil){
					break
				}
				continue
			}
			if(match[i]==')'){
				int2,err = strconv.Atoi(numStr)
				break
			}
			numStr= numStr+ string(match[i])
		}
		sum+=(int1*int2)
	}
	return sum
}
func main(){
	f, err := os.Open("day3/data-input3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	regex1,_ := regexp.Compile(`do\(\)(.*?)don't\(\)`)
	r := bufio.NewReader(f)
	//PART ONE
	// sum:=0
	// for {
	// 	line, errInput := r.ReadString('\n')
	//	sum+= findMulOfMatches(line)
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
	wholeSum:=0
	for {
		line, errInput := r.ReadString('\n')
		firstOccurence :=strings.Index(line,"don't()")
		firstOccurenceOfDo :=strings.Index(line,"do()")

		wholeSum+=findMulOfMatches(line[:firstOccurence])
		result := regex1.FindAllString(line,-1)
		fmt.Println(result)
		for _, res := range result{
			wholeSum+=findMulOfMatches(res)
		}
		lastOccurence := strings.LastIndex(line,"do()")


		if(!strings.Contains(line[lastOccurence:],"don't()")){
			wholeSum+=findMulOfMatches(line[lastOccurence:])
		}
		if errInput != nil {
			break // End of file or read error
		}
	}
	fmt.Println(wholeSum)
}