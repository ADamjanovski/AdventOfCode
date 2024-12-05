package main
import(
	"fmt"
	"os"
	"bufio"
	"strings"
)
func main(){
	sumOfMiddle := 0
	file, err := os.Open("day5/data-input5.txt")
	if err != nil {
		return
	}
	defer file.Close()
	r := bufio.NewReader(file)

	rules :=map[int]int{} 
	for{
		line, err := r.ReadString('\n')
		nums :=strings.Split(line,"|")
		int1,err
		if(err!=nil){
			break
		}
	}
}