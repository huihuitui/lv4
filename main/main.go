package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f := "d:/plan.txt"
	file, err := os.Create(f)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	str := "I'm not afraid of difficulties and insist on learning programming!\n"
	writers := bufio.NewWriter(file)
	_, err3 := writers.WriteString(str)
	if err3 != nil {
		return
	}
	err2 := writers.Flush()
	if err2 != nil {
		return
	}
	//file.Seek(0, os.SEEK_SET)
	readers := bufio.NewReader(file)
	for {
		str, err := readers.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print(str)
				break
			}
			fmt.Println("EOF:", err)
			os.Exit(1)
		}
		fmt.Print(str)
	}
	err1 := file.Close()
	if err1 != nil {
		return
	}
}
