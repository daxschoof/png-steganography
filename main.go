package main

import (
	"flag"
	"fmt"

	"github.com/daxschoof/png-stenography/png"
)

func main() {
	pngPath := flag.String("pngPath", "", "The path to the png you would like to embed")

	flag.Parse()

	original := png.Png {
		Name: *pngPath,
	}

	err := original.ReadFileIntoBinary()
	checkErr(err)

	fmt.Println(original.CheckIfFileIsPng())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}