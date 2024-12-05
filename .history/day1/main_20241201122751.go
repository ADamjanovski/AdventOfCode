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
	var slice1[]int
	var slice2[]int

	for {
		line,errInput:= r.ReadString('\n')
		arr:=strings.Fields(line)
		str1 := arr[0]
		str2 := arr[1]
		int1,err :=strconv.Atoi(str1)
		int2,err := strconv.Atoi(str2)
		if (err!=nil){
			break
		}
		slice1= append(slice1,int1)
		slice2 =append(slice2,int2)
		if errInput != nil {
			fmt.Println(line)
            fmt.Println(errInput)
            break
        }

	}
	sum:=0.0

	slices.Sort(slice1)
	slices.Sort(slice2)

	for i := 0; i < len(slice1); i++ { 
		sum+=math.Abs(float64(slice1[i]-slice2[i]))
	}
	
	integer:=int(sum)
	fmt.Println(integer) 
	//PART TWO
	appearences:= map[int]int{}
	
	
}