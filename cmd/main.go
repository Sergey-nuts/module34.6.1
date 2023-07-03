package main

import (
	"fmt"
	"log"
	"os"

	"module34.6.1/pkg/repkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: regexapp inputfile outputfile")
		return
	}
	inFile := os.Args[1]
	outFile := os.Args[2]

	in, err := os.OpenFile(inFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	repkg.CalcScan(in, out)
}
