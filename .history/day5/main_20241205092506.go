package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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
		if(er!=nil){
			break
		}
		if(err!=nil){
			break
		}
	}

	
	for{
		line, err := r.ReadString('\n')
		line=strings.TrimSpace(line)
		nums :=strings.Split(line,",")

		isGood:=true
		isBad:=false
		appearences:=map[int]int{} 
		middleNum:=0
		if(len(nums)==0){
			break
		}
		for i:=0;i<len(nums);i++{
			num1,err1:=strconv.Atoi(nums[i])
			if(err1!=nil){
				break
			}
			if(i==(len(nums)/2)){
				middleNum+=num1
			}
			for j:=0;j<len(rules[num1]);j++{
				if _, ok := appearences[rules[num1][j]]; ok{			
					isGood=false
					
					continue
				}
			}

			if(isGood){
				appearences[num1]=1
			}else{
				continue
			}
		}
		if(isGood){
			sumOfMiddle+=middleNum
		}
		//PART TWO
		appearences=map[int]int{} 
		var slices []int
		for i:=0;i<len(nums);i++{
			num1,err1:=strconv.Atoi(nums[i])
			slices=append(slices,num1)
			min:=math.MaxInt64
			if(err1!=nil){
				break
			}
			if(i==(len(nums)/2)){
				middleNum+=num1
			}
			for j:=0;j<len(rules[num1]);j++{
				if _, ok := appearences[rules[num1][j]]; ok{			
					isBad=true
					if(appearences[rules[num1][j]]<min){
						min=appearences[rules[num1][j]]
					}
				}
			}
			if(isBad){
				newIndex:=appearences[slices[min]]	
				appearences[slices[min]]=i
				appearences[num1]=newIndex
				slices=append(slices,num1)
			}
		}


		if(err!=nil){
			break
		}
	}
	fmt.Println(sumOfMiddle)
}