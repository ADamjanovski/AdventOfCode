package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findMulOfMatches(result string) int {
	sum := 0
	regex2, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	matches := regex2.FindAllString(result, -1)
	for _, match := range matches {
		var numStr string
		var int1 int
		var int2 int
		var err error
		for i := 4; i < len(match); i++ {
			if match[i] == ',' {
				int1, err = strconv.Atoi(numStr)
				numStr = ""
				if err != nil {
					break
				}
				continue
			}
			if match[i] == ')' {
				int2, err = strconv.Atoi(numStr)
				break
			}
			numStr = numStr + string(match[i])
		}
		sum += (int1 * int2)
	}
	return sum
}
func main() {
	regex1, _ := regexp.Compile(`do\(\)(.*?)don't\(\)`)
	//PART ONE
	// sum:=0
	// for {
	// 	stringData, errInput := r.ReadString('\n')
	//	sum+= findMulOfMatches(stringData)
	// 	matches := regex2.FindAllString(stringData,-1)
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
	wholeSum := 0
	file, err := os.Open("day3/data-input3.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// Read the whole file
	data, err := io.ReadAll(file)
	if err != nil {
	}
	stringData := string(data)
	stopOccurance := strings.Index(stringData, "don't()")
	startOccurance := strings.Index(stringData, "do()")
	if stopOccurance < startOccurance {
		fmt.Println(stringData[:stopOccurance])
		wholeSum += findMulOfMatches(stringData[:stopOccurance])
	}
	result := regex1.FindAllString(stringData, -1)
	for _, res := range result {
		wholeSum += findMulOfMatches(res)
	}
	lastOccurence := strings.LastIndex(stringData, "do()")
	lastOccurenceOfDont := strings.LastIndex(stringData, "don't()")
	if lastOccurence > lastOccurenceOfDont {
		fmt.Println(stringData[lastOccurence:])
		wholeSum += findMulOfMatches(stringData[lastOccurence:])
	}

	fmt.Println(wholeSum)
}
