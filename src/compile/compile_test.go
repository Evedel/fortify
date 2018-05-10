package compile

import(
	"testing"

	"io/ioutil"
	"strings"
)

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

	nTests := 0
	nFails := 0
	nOk := 0

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
					dictionary.Variables = make(map[string]int)

					descr := ""
					inputScr := ""
					inputTex := ""
					inputF90 := ""

					skipTest := false
					if (lines[i][0] == '@') {
						skipTest = true
						i++
						i++
						for lines[i] != "}T{" {
							i++
						}
						i++
						for lines[i] != "}L{" {
							i++
						}
						i++
						for lines[i] != "}F{" {
							i++
						}
						i++
						for lines[i] != "}" {
							i++
						}
						i++
						i++
					} else {
						i++
						for lines[i] != "}T{" {
							descr += strings.TrimSpace(lines[i])
							i++
						}
						i++
						for lines[i] != "}L{" {
							inputScr += strings.TrimSpace(lines[i]) + "\n"
							i++
						}
						i++
						for lines[i] != "}F{" {
							inputTex += strings.TrimSpace(lines[i]) + "\n"
							i++
						}
						inputTex = inputTex[:len(inputTex)-1]
						i++
						for lines[i] != "}" {
							inputF90 += strings.TrimSpace(lines[i]) + "\n"
							i++
						}
						inputF90 = inputF90[:len(inputF90)-1]
						i++
						i++
					}

					if !skipTest {
						tokenised := lexer.Tokenise(inputScr)
						stree, ok, em := syntaxer.BuildTree(tokenised)
						resultTex := toLatex(stree)
						// dictionary.PrintSyntaxTree(stree, "")
						rtree, _ := reduceToFortran(stree)
						// dictionary.PrintSyntaxTree(rtree, "")
						resultF90, okf90 := toFortran(rtree)
						// say.L0("", resultF90, "\n")
						// say.L0("", ok, "\n")
						// say.L0("", okf90, "\n")
						if ok != dictionary.Ok {
							resultTex = dictionary.ErrorCodeDefinitions[ok]
							resultF90 = dictionary.ErrorCodeDefinitions[ok]
						} else if okf90 != dictionary.Ok {
							resultTex = dictionary.ErrorCodeDefinitions[okf90]
							resultF90 = dictionary.ErrorCodeDefinitions[okf90]
						}

						resultF90 = strings.Replace(resultF90, "\t", "", -1)
						nTests += 1
						if (ok != dictionary.Ok) && (verbouse == 1) {
							say.L3(em, "", "\n")
						}
						resTexLines := strings.Split(string(resultTex), "\n")
						if len(resTexLines) > 10 {
							resultTex = strings.Join(resTexLines[6:len(resTexLines)-4], "\n")
						}
						if resultTex == "&\\" {
							resultTex = "-"
						}
						resF90Lines := strings.Split(string(resultF90), "\n")
						if len(resF90Lines) > 2 {
							resultF90 = strings.Join(resF90Lines[1:len(resF90Lines)-1], "\n")
						}
						if len(resF90Lines) == 2 {
							resultF90 = "-"
						}

						if inputTex != resultTex {
							say.L3("#", done, " : fail : ["+f.Name()+"]<->[ " + descr + " ]\n")
							t.Error("For input: [ " + inputScr + " ] Latex:\n Expected [" + inputTex + "]\n Got [" + resultTex + "]")
							nFails += 1
							// t.FailNow()
						} else if inputF90 != resultF90 {
							say.L3("#", done, " : fail : ["+f.Name()+"]<->[ " + descr + " ]\n")
							t.Error("For input: [ " + inputScr + " ] Fortran:\n Expected [" + inputF90 + "]\n Got [" + resultF90 + "]")
							nFails += 1
							// t.FailNow()
						} else {
							say.L1("#", done, " :   ok : ["+f.Name()+"]<->[ " + descr + " ]\n")
							done += 1
							nOk += 1
						}
					}
				}
			}
		}
	}
	say.L2("Tests: ", nTests, "\n")
	say.L2("Fails: ", nFails, "\n")
	say.L2(" Done: ", nOk, "\n")
}
