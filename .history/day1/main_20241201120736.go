package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main(){
	f, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
	r:= bufio.NewReader(f)
	var slice1[]int
	var slice2[]int

	for {
		line,err:= r.ReadString("\n")
		if (err!=nil){
			fmt.Println(err)
			break
		}
		arr:=strings.Fields(line)
		str1 := arr[0]
		str2 := arr[1]
		int1,err :=strconv.Atoi(str1)
		int2,err := strconv.Atoi(str2)
		slice1= append(slice1,int1)
		slice2 =append(slice2,int2)
		var sum:=0
		for i := 0; i < len(slice1); i++ { 
			sum+=math.Abs(float64(slice1[i]-slice2[i]))
		} 
	}

}