package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println(s)
		for i := 1; i < len(arr); i++ { 
			int1,err:=strconv.Atoi(arr[i])		
			slice= append(slice,int1)
			if(-3<(slice[i]-slice[i-1]) && (slice[i]-slice[i-1])<(-1)){
				if(i==1){
					decreasing=true
				}else{
					if(!decreasing){
						safe=false
						break
					}
				}
			}else if(3>(slice[i]-slice[i-1]) && (slice[i]-slice[i-1])>(-1)){
				if(i==1){
					decreasing=false
				}else{
					if(decreasing){
						safe=false
						break
					}
				}
			}else{
				safe=false
				break
			}
			if (err!=nil){
				break
			}
		}
		
		if(safe){
			safeLevels++
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
	fmt.Println(safeLevels) 
	// //PART TWO
	// appearences:= map[int]int{}
	// for i := 0; i < len(slice1); i++ { 
	// 	appearences[slice1[i]]=0
	// }
	// for i := 0; i < len(slice1); i++ {

	// 	appearences[slice2[i]]= appearences[slice2[i]]+1
	// }
	// similarity_score:=0
	// for i := 0; i < len(slice1); i++ {

	// 	similarity_score+= (appearences[slice1[i]]*slice1[i])
	// }
	// fmt.Println(similarity_score) 

}