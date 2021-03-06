package main

import(
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

import(
	"github.com/Evedel/fortify/src/say"
	"github.com/Evedel/fortify/src/lexer"
	"github.com/Evedel/fortify/src/compile"
	"github.com/Evedel/fortify/src/syntaxer"
	"github.com/Evedel/fortify/src/dictionary"
)

func main() {
	say.Init("3")
	dictionary.Init()

	source := ""
	flag.Parse()
	tailInput := flag.Args()
	say.L1("Input File: ", tailInput, "\n")

	if len(tailInput) == 1 {
		sourceInputFull := tailInput[0]
		sourceName := [3]string{sourceInputFull,
			sourceInputFull,
			""}
		idot := strings.LastIndex(sourceInputFull, ".")
		isls := strings.LastIndex(sourceInputFull, "/")
		if idot != -1 {
			if isls != -1 {
				if isls < idot {
					// adfasdfasf/asfaf.asdf
					sourceName[0] = sourceInputFull[isls+1 : idot]
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
		say.L1("Working Directory: ", sourceName, "\n")
		if _, err := os.Stat(sourceInputFull); os.IsNotExist(err) {
			say.L3("File doesn't exist:", sourceName, "\n")
		} else {
			err := os.Mkdir(sourceName[2]+"/"+sourceName[0], 0744)
			if (err == nil) || (os.IsExist(err)) {
				if b, err := ioutil.ReadFile(sourceInputFull); err != nil {
					fmt.Print(err)
				} else {
					source = source + string(b)
				}
				tokenisedForm := lexer.Tokenise(source)
				// say.L1("", tokenisedForm, "\n")
				SyntaxTree, ok, emf := syntaxer.BuildTree(tokenisedForm)
				dictionary.PrintSyntaxTree(SyntaxTree, "")
				if ok == dictionary.Ok {
					// dictionary.PrintSyntaxTree(SyntaxTree, "")
					compile.ToLaTeX(SyntaxTree, sourceName)
					compile.ToFortran(SyntaxTree, sourceName)
					// compile.ToClang(SyntaxTree, sourceName)
				} else {
					say.L3(emf, "", "\n")
				}
			} else {
				say.L3("Error:", err, "\n")
			}
		}
	}
	say.L1("Done. ", "", "\n")
}
