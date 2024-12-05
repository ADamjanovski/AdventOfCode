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
	r:= f.re
}