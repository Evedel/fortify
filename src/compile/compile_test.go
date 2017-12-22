package compile

import(
  "testing"

  "io/ioutil"
  "strings"

  "say"
  "lexer"
  "syntaxer"
  "dictionary"
)

func TestExpression(t *testing.T) {
  say.Init("3")
  dictionary.Init()
  verbouse := 0

  prefix := "../../_testdata"
  tests, err := ioutil.ReadDir(prefix)
  if err != nil {
      say.L3("", err, "\n")
  } else {
    for _, f := range tests {
      content, err := ioutil.ReadFile(prefix+"/"+f.Name())
      if err != nil {
      	say.L3("", err, "\n")
      } else {
        lines := strings.Split(string(content), "\n")
        i := 0
        for i < len(lines) {
          dictionary.Variables = make(map[string]int)
          teststr := lines[i]
          lres := lines[i+1]
          fres := lines[i+2]
          i = i + 4
          tokenised := lexer.Tokenise(teststr + "\n")
          stree, ok, em := syntaxer.BuildTree(tokenised)
          // say.L2(em, "", "\n")
          lsrc := toLatex(stree)
          fsrc := toFortran(stree)
          if lsrc == "" { lsrc = "-"}
          if fsrc == "" { fsrc = "-"}
          if ok == false { lsrc = "-1"; fsrc = "-1"}
          fsrc = strings.Replace(fsrc, "\t", "", -1)
          fsrc = strings.Replace(fsrc, "\n", "", -1)
          lsrc = strings.Replace(lsrc, "\n", "", -1)
          dictionary.PrintSyntaxTree(stree, "")
          if (! ok) && (verbouse==1) { say.L3(em, "", "\n")}
          if (lsrc != lres) {
            say.L1("[" + teststr + "] <-> fail", "", "\n")
            t.Error("For input: [" + teststr + "]\n Latex src expected to be [" + lres + "]\n Got [" + lsrc + "]")
            t.FailNow()
          } else if (fsrc != fres) {
            say.L1("[" + teststr + "] <-> fail", "", "\n")
            t.Error("For input: [" + teststr + "]\n Fortran src expected to be [" + fres + "]\n Got [" + fsrc + "]")
            t.FailNow()
          } else {
            say.L1("[" + teststr + "] <-> ok", "", "\n")
          }
        }
      }
    }
  }
}
