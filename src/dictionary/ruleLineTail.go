package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func RuleCommentTex(ttail []Token) (correct bool, stopInd int, childs []TokenNode, errmsg string) {
  lentt := len(ttail)
  correct = false
  stopInd = 0
  errmsg = "Default tex comment form form is: % any expression that will be compiled in f90.\\n"
  ich := 1
  childs = append(childs, TokenNode{ttail[0], nil})
  for ich < lentt {
    childs = append(childs, TokenNode{ttail[ich], nil})
    if ttail[ich].Id == DoubleQuote {
      correct = true
      stopInd = ich
      errmsg = ""
      return
    }
    ich += 1
  }
  return
}
