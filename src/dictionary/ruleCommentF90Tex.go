package dictionary

// import (
//   "say"
// )

func ruleCommentF90Tex(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
  indexInternal := 0
  indexChild := 0
  chchilds := []TokenNode{}
  for indexInternal < len(ttail) {
    if ttail[indexInternal].Id == CarriageReturn {
      childs = append(
        childs,
        TokenNode{
          Token{ Expression, "expression", ""},
          append([]TokenNode{}, TokenNode{Token{ CarriageReturn, "\\n", ""}, nil})})
      resCode = Ok
      stopInd = indexInternal
      errmsg = ""
      return
    } else {
      resCode, indexChild, chchilds, errmsg = RuleExpression(ttail[indexInternal:])
      if resCode == Ok {
        childs = append(
          childs,
          TokenNode{Token{ Expression, "expression", ""}, chchilds})
        indexInternal += indexChild
      } else {
        return
      }
    }
    indexInternal += 1
  }
  return
}
