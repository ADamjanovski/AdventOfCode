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
	f, err := os.Open("day1/data-input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
	r:= bufio.NewReader(f)
	safeLevels:=0
	for {
		line,errInput:= r.ReadString('\n')
		arr:=strings.Fields(line)
		safe := true
		var slice[]int
		decreasing:=true
		int1,err:=strconv.Atoi(arr[0])		
		slice= append(slice,int1)
		for i := 1; i < len(arr); i++ { 
			int1,err:=strconv.Atoi(arr[i])		
			slice= append(slice,int1)
			if((slice[i]-slice[i-1])<-0){
				
			}
			if (err!=nil){
				break
			}
		}
		
		if (err!=nil){
			break
		}
		if errInput != nil {
			fmt.Println(line)
            fmt.Println(errInput)
            break
        }

	}
	sum:=0.0

	for i := 0; i < len(slice1); i++ { 
		sum+=math.Abs(float64(slice1[i]-slice2[i]))
	}
	
	integer:=int(sum)
	fmt.Println(integer) 
	//PART TWO
	appearences:= map[int]int{}
	for i := 0; i < len(slice1); i++ { 
		appearences[slice1[i]]=0
	}
	for i := 0; i < len(slice1); i++ {

		appearences[slice2[i]]= appearences[slice2[i]]+1
	}
	similarity_score:=0
	for i := 0; i < len(slice1); i++ {

		similarity_score+= (appearences[slice1[i]]*slice1[i])
	}
	fmt.Println(similarity_score) 

}