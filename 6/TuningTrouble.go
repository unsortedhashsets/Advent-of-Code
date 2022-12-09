package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkBuffer(bufferSize int, line string){
	for i := range line {
		buffer := make(map[byte]bool)
		for j:=0; j<bufferSize; j++{
			buffer[line[i+j]]=true
		}
		if len(buffer) == bufferSize{
			fmt.Println(i+bufferSize)
			break
		}
	}
}

func main(){
	
	readFile, err := os.Open("input.txt")
	
    if err != nil {
        fmt.Println(err)
    }

	fs := bufio.NewScanner(readFile)
	fs.Scan()
	
	checkBuffer(4, fs.Text())
	checkBuffer(14, fs.Text())
	
	readFile.Close()
}