package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func diagonalBackwards(strMatrix[][]string,i int,j int) int{
	count:=0
	if(strMatrix[i-1][j-1]=="M" && strMatrix[i-2][j-2]=="A" && strMatrix[i-3][j-3]=="S"){count++}
	return count
}
func diagonalForwards(strMatrix[][]string,i int,j int) int{
	count:=0
	if(strMatrix[i+1][j+1]=="M" && strMatrix[i+2][j+2]=="A" && strMatrix[i+3][j+3]=="S"){count++}
	return count
}
func diagonalUpright(strMatrix[][]string,i int,j int) int{
	count:=0
	if(strMatrix[i-1][j+1]=="M" && strMatrix[i-2][j+2]=="A" && strMatrix[i-3][j+3]=="S"){count++}
	return count
}
func diagonalDownleft(strMatrix[][]string,i int,j int) int{
	count:=0
	if(strMatrix[i+1][j-1]=="M" && strMatrix[i+2][j-2]=="A" && strMatrix[i+3][j-3]=="S"){count++}
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

	var charsMatrix [][] string
	for{
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		chars :=strings.Split(line,"")
		charsMatrix = append(charsMatrix, chars)
		if(err!=nil){
			break
		}
	}

	for i:=0; i<len(charsMatrix);i++{
		for j:=0; j<len(charsMatrix[i]);j++{
			if(i>=3 && j>=3 && charsMatrix[i][j]=="X"){
				wholeSum+=diagonalBackwards(charsMatrix,i,j)
			}
			if(i>=3 && (len(charsMatrix[i])-j)>3 && charsMatrix[i][j]=="X"){
				wholeSum+=diagonalUpright(charsMatrix,i,j)
			}
			if((len(charsMatrix)-i)>3  && j>=3 && charsMatrix[i][j]=="X"){
				wholeSum+=diagonalDownleft(charsMatrix,i,j)
			}
			if((len(charsMatrix)-i)>3 && (len(charsMatrix[i])-j)>3 && charsMatrix[i][j]=="X"){
				wholeSum+=diagonalForwards(charsMatrix,i,j)
			}
			if(charsMatrix[i][j]=="X" && j>=3){
				if(charsMatrix[i][j-1]=="M" && charsMatrix[i][j-2]=="A" && charsMatrix[i][j-3]=="S"){
					wholeSum++
				}
			}
			if(charsMatrix[i][j]=="X" && (len(charsMatrix[i])-j)>3){
				if(charsMatrix[i][j+1]=="M" && charsMatrix[i][j+2]=="A" && charsMatrix[i][j+3]=="S"){
					wholeSum++
				}
			}
			if(i>=3 && charsMatrix[i][j]=="X"){
				if(charsMatrix[i-1][j]=="M" && charsMatrix[i-2][j]=="A" && charsMatrix[i-3][j]=="S"){
					wholeSum++
				}
			}
			if((len(charsMatrix)-i)>3 && charsMatrix[i][j]=="X"){
				if(charsMatrix[i+1][j]=="M" && charsMatrix[i+2][j]=="A" && charsMatrix[i+3][j]=="S"){
					wholeSum++
				}
			}
		}
	}
	fmt.Println(wholeSum)

	//PART TWO
	for i:=0; i<len(charsMatrix);i++{
		for j:=0; j<len(charsMatrix[i]);j++{
			if(charsMatrix[i][j]=="A" && i>=1 && j>=1 && (len(charsMatrix)-i)>1 && (len(charsMatrix[i])))
		}
	}
	
}