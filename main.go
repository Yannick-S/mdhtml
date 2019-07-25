package main

import (
	"fmt"
	"flag"
	"errors"
	"os"
	"io/ioutil"
	"./mdhtml"
)

func main() {
	//Define the flags
	inFilePtr := flag.String("in", "", "input file name")
	outFilePtr := flag.String("out", "", "output file name")

	//Parse the flags
	flag.Parse()
	inFile := *inFilePtr
	outFile := *outFilePtr

	// after this block, b is the input file
	var b []byte
	if inFile == ""{
		//TODO: Implement the following
		fmt.Println("Pipeing should be here, not implemented yet")
		panic(errors.New("Not Implemented"))
	} else {
		//Open File for Reading
		file, err := os.Open(inFile)
    	if err != nil {
        	panic(err)
		}
		defer file.Close()
	
		b, err = ioutil.ReadAll(file)
		if err != nil {
        	panic(err)
		}
	}

	//Conversion happens here


	blocks := mdhtml.Convert(b)
	for _, block := range blocks{
		fmt.Println("New Block:")
		for _, line := range block {
			fmt.Println("\tLine:", line)
		}
	}

	//TODO: Implement the following
	if outFile == ""{
		fmt.Println("Saving to standard file should be implemented here, not implemented yet")
		panic(errors.New("Not Implemented"))
	} else {
		fmt.Println("Implemnt saving to file here")
	}
}
