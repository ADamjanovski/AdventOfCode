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

	rules :=map[int][]int{} 
	for{
		line, err := r.ReadString('\n')
		line=strings.TrimSpace(line)
		nums :=strings.Split(line,"|")
		if(len(nums)==1){
			break
		}
		int1,er:=strconv.Atoi(nums[0])
		int2,er:=strconv.Atoi(nums[1])
		rules[int1]=append(rules[int1], int2)
		fmt.Println(int1)
		fmt.Println(int2)
		fmt.Println("end of line")
		if(er!=nil){
			break
		}
		if(err!=nil){
			break
		}
	}

	appearences:=map[int]int{} 
	fmt.Println(rules)
	for{
		line, err := r.ReadString('\n')
		nums :=strings.Split(line,",")
		isGood:=true
		middleNum:=0
		if(len(nums)==0){
			break
		}
		for i:=0;i<len(nums);i++{
			num1,err1:=strconv.Atoi(nums[i])
			if(err1!=nil){
				break
			}
			if(i==(len(nums)+1)){
				middleNum+=num1
			}
			for j:=0;j<len(rules[num1]);j++{
				if _, ok := appearences[rules[num1][j]]; ok{			
					isGood=false
					continue
				}
			}
			if(isGood){
				a
			}
		}
		if(isGood){
			sumOfMiddle+=middleNum
		}
		if(err!=nil){
			break
		}
	}
	fmt.Println(sumOfMiddle)
}