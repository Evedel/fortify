package dictionary

// import (
//   "say"
// )

func RuleExpression(ttail []Token) (correct bool, stopInd int, childs []TokenNode, errmsg string) {
  correct = false
  stopInd = 0
  errmsg = "Not a correct expression."
  indx := 0
  for indx < len(ttail) {
    // say.L0("", ttail[indx], "\n")
    if ttail[indx].Id == Space {
      // ignore
    } else if ttail[indx].Id == CarriageReturn {
      correct = true
      stopInd = indx
      errmsg = ""
      childs = append([]TokenNode{}, TokenNode{Token{ CarriageReturn, "\\n", 0, 0, ""}, nil})
      return
    } else if ttail[indx].Id == Print {
      ok, is, chch, erms := RulePrint(ttail[indx+1:])
      correct = ok
      if ok {
        stopInd = indx + is + 1
        childs = append([]TokenNode{}, TokenNode{ttail[indx], chch})
      } else {
        errmsg = erms
      }
      return
    // } else if ttail[indx].Id == CommentTex {
    //
    // } else if ttail[indx].Id == CommentF90 {
    //
    } else {
      errmsg = "Cannot use " + ttail[indx].IdName + " in expression. " + errmsg
      return
    }
    indx += 1
  }
  return
}
