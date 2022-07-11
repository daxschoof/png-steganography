package main

import (
	"bufio"
	"fmt"
	"os"
)

type png struct {
	name string
	byteArr []byte
}

func main() {
	original := png {
		name: "test.png",
	}

	err := original.readFileIntoBinary()
	checkErr(err)

	fmt.Println(original.checkIfFileIsPng())
}

func (p *png) readFileIntoBinary() error {
	file, err := os.Open(p.name)
	if err != nil {
		return err
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return err
	}

 	fileSize := stats.Size()
	byteArr := make([]byte, fileSize)

	buf := bufio.NewReader(file)
	_, err = buf.Read(byteArr)
	if err != nil {
		return err
	}

	p.byteArr = byteArr
	
	return nil
}

func (p *png) checkIfFileIsPng() bool {
	if len(p.byteArr) < 8 {
		return false
	}

	firstEightPng := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	return checkSliceEquality(firstEightPng, p.byteArr[0:8])
}

func checkSliceEquality [K comparable] (arr1, arr2 []K) bool {
	if arr1 == nil || arr2 == nil || len(arr1) != len(arr2){
		return false 
	}
	
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}