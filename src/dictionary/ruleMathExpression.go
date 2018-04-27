package dictionary

import (
  "github.com/Evedel/fortify/src/say"
)

func ruleMathExpression(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
  thisName := "Math Expression: "
	resCode = UndefinedError
	stopInd = 0
	index := 0
	chindex := 0
  resNode = TokenNodeOperand()

  for index < len(ttail) {
    say.L2("", ttail[index], "\n")
    tokenid := ttail[index].Id
    tokenvalstr := ttail[index].Value
    if tokenid == Space {
      resNode.List = append(resNode.List, TokenNodeSpace())
    } else if tokenid == Word {
      if _, ok := Variables[tokenvalstr]; ok {
        // resCode, chindex, resNode, errmsg = ruleOperator(ttail[index:])
        stopInd = index + chindex
      } else {
        resCode = NotALanguageKeyWord
        errmsg = thisName + "There is no defined rule for [ " + tokenvalstr + " ] symbol."
        return
      }
    } else if tokenid == CarriageReturn {
      resNode.List = append(resNode.List, TokenNodeReturn())
    } else {
      resCode = NotALanguageKeyWord
      errmsg = thisName + "There is no defined rule for [ " + ttail[index].IdName + " ] symbol."
      return
    }
  }
  return
}
