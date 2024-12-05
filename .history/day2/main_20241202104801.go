package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func  isSafe( int1 int, int2 int, decreasing bool,i int)(bool,bool){
	if(-3<=(int1-int2) && (int1-int2)<=(-1)){
		if(i==1){
			decreasing=true
		}else{
			if(!decreasing){
				return false,decreasing
			}
		}
	}else if(3>=(int1-int2) && (int1-int2)>=(1)){
		if(i==1){
			decreasing=false
		}else{
			if(decreasing){
				return false,decreasing
			}
		}
	}
	return true, decreasing
}
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
		removedLevel:=false
		int1,err:=strconv.Atoi(arr[0])		
		slice= append(slice,int1)
		for i := 1; i < len(arr); i++ { 
			int1,err:=strconv.Atoi(arr[i])
			safe,decreasing=isSafe(int1,slice[len(slice)-1],decreasing,i)	
			if(safe){
				slice=append(slice,int1)
			}else{
				i++
				int2,err:=strconv.Atoi(arr[i])
				safe,decreasing=isSafe(int2,slice[len(slice)-1],decreasing,i)	
				if(safe){
					slice = append(slice, int2)
				}
				safe,decreasing=isSafe(int1,int2,decreasing,i)	

				//PART TWO ADDING REMOVAL OF ONE SAFE LEVEL
				if(!removedLevel){

					removedLevel=true
					continue
				}
				fmt.Println(slice)	
				fmt.Println(int1)
				fmt.Println(arr)
				fmt.Println("END OF LINE")
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
            fmt.Println(errInput)
            break
        }

	}
	fmt.Println(safeLevels) 
//PART TWO
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