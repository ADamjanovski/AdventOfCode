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

	var rules
	for{
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		chars :=strings.Split(line,"")
		charsMatrix = append(charsMatrix, chars)
		if(err!=nil){
			break
		}
	}
}