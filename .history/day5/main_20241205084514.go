package main
import(
	"fmt"
	"os"
	"bufio"
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
		line = strings.TrimSpace(line)
		chars :=strings.Split(line,"|")
		
		if(err!=nil){
			break
		}
	}
}