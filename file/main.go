package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	fileWrite, err1 := os.OpenFile("./output.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, os.FileMode(0777))
	defer fileWrite.Close()
	if err1 != nil {
		fmt.Println("error is: ", err1)
	}
	if err != nil {
		fmt.Println("error is: ", err)
		return
	}
	defer file.Close()
	temp := make([]byte, 128)
	var content []byte
	for {
		size, err := file.Read(temp)
		if err == io.EOF {
			//fmt.Println("eof")
			break
		}
		if err != nil {
			fmt.Println("error is: ", err)
			return
		}
		fileWrite.Write(temp[:size])
		content = append(content, temp[:size]...)

	}
	fmt.Println(string(content))
	/* fileWrite, err1 := os.OpenFile("./output.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(0777))
	if err1 != nil {
		fmt.Println("error is: ", err1)
	}
	defer fileWrite.Close()
	fileWrite.Write(content) */
}
