package main

func main(){
	sumOfMiddle := 0
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
}