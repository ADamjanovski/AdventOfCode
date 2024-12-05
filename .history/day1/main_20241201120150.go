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
		 arr:=[...]string{strings.Fields(line)}
		
		append(slice1,strconv.Atoi(str1))
		append(slice2,strconv.Atoi(str2))
	}

}