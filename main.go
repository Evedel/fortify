package main

import (
	"fmt"
	"flag"
	"io/ioutil"

	"compile"
)

func main() {
	source := ""
	flag.Parse()
	tailInput := flag.Args()
	fmt.Println("-> Input File: ", tailInput)
	if len(tailInput) == 1 {
		if b, err := ioutil.ReadFile(tailInput[0]); err != nil {
				fmt.Print(err)
		} else {
			source = source + string(b)
		}
		compile.LaTeX(source, tailInput[0])
	}
	fmt.Printf("-> Done\n")
}
