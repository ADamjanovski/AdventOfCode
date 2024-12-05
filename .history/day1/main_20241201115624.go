package main
import (    "bufio"
"fmt"
"os"
)

func main(){
	f, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
	r:= bufio.NewReader(f)
	var array1:= []int

	for {
		line,err:= r.ReadString("\n")
		if (err!=nil){
			fmt.Println(err)
			break
		}

	}
}