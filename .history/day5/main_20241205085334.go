package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main(){
	sumOfMiddle := 0
	file, err := os.Open("day5/data-input5.txt")
	if err != nil {
		return
	}
	defer file.Close()
	r := bufio.NewReader(file)

	rules :=map[int]int{} 
	for{
		line, err := r.ReadString('\n')
		nums :=strings.Split(line,"|")
		if(len(nums)==0){
			break
		}
		int1,er:=strconv.Atoi(nums[0])
		int2,er:=strconv.Atoi(nums[1])
		rules[int1]=int2
		if(er!=nil){
			break
		}
		if(err!=nil){
			break
		}
	}

	appearences:=map[int]int{} 
	for{
		line, err := r.ReadString('\n')
		nums :=strings.Split(line,",")
		if(len(nums)==0){
			break
		}
		for i:=0;i<len(nums);i++{
			num1,err1:=strconv.Atoi(nums[i])
			if(err1!=nil){
				break
			}
			num2:=rules[num1]
			if_, ok := myMap[key]; ok{
				continue
			}
		}
		
		if(err!=nil){
			break
		}
	}
}