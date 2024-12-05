package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func  isSafe( int1 int, int2 int, decreasing bool,i int)(bool,bool){
	if(-3<=(int2-int1) && (int2-int1)<=(-1)){
		if(i==1){
			decreasing=true
		}else{
			if(!decreasing){
				return false,decreasing
			}
		}
	}else if(3>=(int2-int1) && (int2-int1)>=(1)){
		if(i==1){
			decreasing=false
		}else{
			if(decreasing){
				return false,decreasing
			}
		}
	}else{
		return false, decreasing
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
		var decreasing bool
		removedLevel:=false
		int1,err:=strconv.Atoi(arr[0])		
		slice= append(slice,int1)
		for i := 1; i < len(arr); i++ { 
			int1,err:=strconv.Atoi(arr[i])
			safe,decreasing=isSafe(slice[len(slice)-1],int1,decreasing,len(slice))	
			if(safe){
				slice=append(slice,int1)
			}else{
				if(removedLevel){
					safe=false
					break
				}
				i++
				if(i==len(arr)&& !removedLevel){
					safe=true
					break
				}
				int2,err:=strconv.Atoi(arr[i])
				if (err!=nil){
					fmt.Println(err)
					break
				}
				safe,decreasing=isSafe(slice[len(slice)-1],int2,decreasing,len(slice))	
				if(safe){
					slice = append(slice, int2)
					removedLevel=true
					continue
				}else{
					if(len(slice)>1){
						safe,decreasing=isSafe(slice[len(slice)-2],int1,decreasing,len(slice)-1)
						if(safe){
							safe,decreasing=isSafe(int1,int2,decreasing,len(slice))
							if(safe){
								slice[len(slice)-1]=int1
								slice = append(slice, int2)
								removedLevel=true
								continue
							}else{
								safe=false
								break
							}
						}else{
							safe=false
							break
						}	
						}else if(len(slice)==1){

					}
					if(safe || ){
						safe,decreasing=isSafe(int1,int2,decreasing,len(slice))
						if(safe){
							slice[len(slice)-1]=int1
							slice = append(slice, int2)
							removedLevel=true
							continue
						}else{
							safe=false
							break
						}
					}else{
						safe=false
						break
					}	
				}
				}

				//PART TWO ADDING REMOVAL OF ONE LEVEL
			}
			if (err!=nil){
				break
			}
		}
		
		if(safe){
			if(len(slice)!=len(arr)){
				fmt.Println(slice)
				fmt.Println(arr)
				fmt.Println("END OF LINE")
			}
			safeLevels++
		}
		if (err!=nil){
			break
		}
		if errInput != nil {
			fmt.Println(slice)
            fmt.Println(errInput)
            break
        }

	}
	fmt.Println(safeLevels) 

}