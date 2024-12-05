package main
import ("fmt"

func main(){
	f, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
}