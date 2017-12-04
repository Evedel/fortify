package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"io/ioutil"

	"compile"
)

func main() {
	source := ""
	flag.Parse()
	tailInput := flag.Args()
	fmt.Println("-> Input File: ", tailInput)
	if len(tailInput) == 1 {
		sourceInputFull := tailInput[0]
		sourceName := [3]string{ sourceInputFull,
														sourceInputFull,
														"" }
		idot := strings.LastIndex(sourceInputFull, ".")
		isls := strings.LastIndex(sourceInputFull, "/")
		if idot != -1 {
			if isls != -1 {
				if isls < idot {
					// adfasdfasf/asfaf.asdf
					sourceName[0] = sourceInputFull[isls+1:idot]
					sourceName[1] = sourceInputFull[isls+1:]
					sourceName[2] = sourceInputFull[:isls]
				} else {
					// sdfasdfa.sdfasdf/asdfadsf
					sourceName[0] = sourceInputFull[isls+1:]
					sourceName[1] = sourceInputFull[isls+1:]
					sourceName[2] = sourceInputFull[:isls]
				}
			} else {
				// adsfaf.adfad
				sourceName[0] = sourceInputFull[:idot]
				sourceName[1] = sourceInputFull
				sourceName[2] = ""
			}
		} else {
			if isls != -1 {
				// adfadsf/adfadf
				sourceName[0] = sourceInputFull[isls+1:]
				sourceName[1] = sourceInputFull[isls+1:]
				sourceName[2] = sourceInputFull[:isls]
			} else {
				// adfasdfads
        // default from definition
			}
		}
		fmt.Print("-> Working Directory: ", sourceName, "\n")
		if _, err := os.Stat(sourceInputFull); os.IsNotExist(err) {
			fmt.Printf("File %s doesn't exist.\n", sourceName)
		} else {
			err := os.Mkdir(sourceName[2] + "/" + sourceName[0], 0744)
			if (err == nil) || (os.IsExist(err)) {
				if b, err := ioutil.ReadFile(sourceInputFull); err != nil {
						fmt.Print(err)
				} else {
					source = source + string(b)
				}
				compile.LaTeX(source, sourceName)
			} else {
				fmt.Print("-> Error: ", err, "\n")
			}
		}
	}
	fmt.Printf("-> Done\n")
}
