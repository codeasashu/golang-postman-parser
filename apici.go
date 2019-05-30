package main

import (
	"bufio"
	//	"encoding/json"
	"fmt"
	"os"
)

type requestdata struct {
	Url  string `json:"uri"`
	Host string `json:"host"`
}

func main() {
	if r := len(os.Args); r < 2 {
		panic("Not enough args")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
		os.Exit(2)
	}

	buffer := bufio.NewScanner(file)
	//	buffer := bufio.NewReader(file)
	//	bytt, _ := buffer.ReadBytes('\n')
	//	fmt.Println(*bytt)
	//	fmt.Printf("%T", bytt)
	for buffer.Scan() {
		fmt.Println(buffer.Text())
		os.Exit(1)
	}
	//fmt.Println("Something", &file.Read())
}
