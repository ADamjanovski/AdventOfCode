package main

import (
	"bufio"
	"fmt"
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
		append(slice1int1
		append(slice2,strconv.Atoi(str2))
	}

}