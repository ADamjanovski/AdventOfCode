package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)
func diagonal

func main(){
	wholeSum := 0
	file, err := os.Open("day4/data-input4.txt")
	if err != nil {
		return
	}
	defer file.Close()
	r := bufio.NewReader(file)

	stringData := string(data)
	var charsMatrix [][] string
	for{
		line, err := r.ReadString("\n")
		strings.Split(line,"")
		charsMatrix = append(charsMatrix, line)
		if(err!=nil){
			break
		}
	}
	fmt.Println(charsMatrix)

	for i:=0; i<len(charsMatrix);i++{
		for j:=0; j<len(charsMatrix[i]);i++{

		}
	}
}