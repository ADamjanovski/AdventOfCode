package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)
func diagonal(strMatrix[][]string,i int,j int) int{
	count:=0
	if(strMatrix[i-1][j-1]=="M" && strMatrix[i-2][j-2]=="A" && strMatrix[i-3][j-3]=="S"){count++}
	if(strMatrix[i+1][j+1]=="M" && strMatrix[i+2][j+2]=="A" && strMatrix[i+3][j+3]=="S"){count++}
	return count
}

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
			if(i>=3 && charsMatrix[i][j]=="X"){
				wholeSum+=diagonal(charsMatrix,i,j)
			}
			if(charsMatrix[i][j]=="X" && j>=3){

			}
		}
	}
}