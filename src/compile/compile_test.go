package compile

import (
	"testing"

	"io/ioutil"
	"strings"
)

// import (
// 	"dictionary"
// 	"lexer"
// 	"say"
// 	"syntaxer"
// )

import(
	"github.com/Evedel/fortify/src/dictionary"
	"github.com/Evedel/fortify/src/lexer"
	"github.com/Evedel/fortify/src/say"
	"github.com/Evedel/fortify/src/syntaxer"
)

func TestExpression(t *testing.T) {
	say.Init("3")
	dictionary.Init()
	verbouse := 0
	done := 0

	prefix := "../../_testdata"
	tests, err := ioutil.ReadDir(prefix)
	if err != nil {
		say.L3("", err, "\n")
	} else {
		for _, f := range tests {
			content, err := ioutil.ReadFile(prefix + "/" + f.Name())
			if err != nil {
				say.L3("", err, "\n")
			} else {
				lines := strings.Split(string(content), "\n")
				i := 0
				for i < len(lines) {
					inputScr := ""
					inputTex := ""
					inputF90 := ""
					dictionary.Variables = make(map[string]int)
					i += 1
					for lines[i] != "}L{" {
						inputScr += strings.TrimSpace(lines[i]) + "\n"
						i += 1
					}
					i += 1
					for lines[i] != "}F{" {
						inputTex += strings.TrimSpace(lines[i]) + "\n"
						i += 1
					}
					inputTex = inputTex[:len(inputTex)-1]
					i += 1
					for lines[i] != "}" {
						inputF90 += strings.TrimSpace(lines[i]) + "\n"
						i += 1
					}
					i += 2

					tokenised := lexer.Tokenise(inputScr)
					stree, ok, em := syntaxer.BuildTree(tokenised)
					resultTex := toLatex(stree)
					resultF90 := toFortran(stree)
					if resultTex == "" {
						resultTex = "-"
					}
					if resultF90 == "" {
						resultF90 = "-\n"
					}

					if ok != dictionary.Ok {
						resultTex = dictionary.ErrorCodeDefinitions[ok]
						resultF90 = dictionary.ErrorCodeDefinitions[ok] + "\n"
					}
					resultF90 = strings.Replace(resultF90, "\t", "", -1)
					// fsrc = strings.Replace(fsrc, "\n", "", -1)
					// lsrc = strings.Replace(lsrc, "\n", "", -1)
					// dictionary.PrintSyntaxTree(stree, "")
					if (ok != dictionary.Ok) && (verbouse == 1) {
						say.L3(em, "", "\n")
					}
					if inputTex != resultTex {
						say.L1("#", done, " fail <-> ["+f.Name()+"] \n")
						t.Error("For input: [" + inputScr + "] Latex:\n Expected [" + inputTex + "]\n Got [" + resultTex + "]")
						t.FailNow()
					} else if inputF90 != resultF90 {
						say.L1("#", done, " fail <-> ["+f.Name()+"] \n")
						t.Error("For input: [" + inputScr + "] Fortran:\n Expected [" + inputF90 + "]\n Got [" + resultF90 + "]")
						t.FailNow()
					} else {
						say.L1("#", done, " ok <-> ["+f.Name()+"] \n")
						done += 1
					}
				}
			}
		}
	}
}
