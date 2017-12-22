package dictionary

// import (
//   "say"
// )

func ruleCommentF90Tex(ttail []Token) (correct bool, stopInd int, childs []TokenNode, errmsg string) {
  indexInternal := 0
  for indexInternal < len(ttail) {
    if ttail[indexInternal].Id == CarriageReturn {
      childs = append(
        childs,
        TokenNode{
          Token{ Expression, "expression", 0, 0, ""},
          append([]TokenNode{}, TokenNode{Token{ CarriageReturn, "\\n", 0, 0, ""}, nil})})
      correct = true
      stopInd = indexInternal
      errmsg = ""
      return
    } else {
      ok, is, chch, erms := RuleExpression(ttail[indexInternal:])
      if ok {
        childs = append(childs, TokenNode{Token{ Expression, "expression", 0, 0, ""}, chch})
        indexInternal += is
      } else {
        errmsg = erms
        return
      }
    }
    indexInternal += 1
  }
  return
}
