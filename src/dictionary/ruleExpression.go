package dictionary

// import (
//   "say"
// )

func RuleExpression(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
  resCode = Ok
  stopInd = 0
  index := 0
  chindex := 0
  chchilds := []TokenNode{}
  for index < len(ttail) {
    // say.L1("", ttail[index], "\n")
    if ttail[index].Id == Space {
      stopInd = index
      childs = append([]TokenNode{}, TokenNode{ttail[index], nil})
      return
    } else if ttail[index].Id == CarriageReturn {
      stopInd = index
      childs = append([]TokenNode{}, TokenNode{Token{ CarriageReturn, "\\n", 0, 0, ""}, nil})
      return
    } else if ttail[index].Id == Print {
      resCode, chindex, chchilds, errmsg = RulePrint(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if (ttail[index].Id == CommentTex) || (ttail[index].Id == CommentF90) {
      resCode, chindex, chchilds, errmsg = ruleCommentF90Tex(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if ttail[index].Id == CommentAll {
      resCode, chindex, chchilds, errmsg = ruleCommentAll(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if ttail[index].Id == DeclarationVar {
      resCode, chindex, chchilds, errmsg = ruleDeclaration(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else {
      resCode = NotALanguageKeyWord
      if ttail[index].Id == Word {
        errmsg = "There is no defined rule for [ " + ttail[index].ValueStr + " ] symbol."
      } else {
        errmsg = "There is no defined rule for [ " + ttail[index].IdName + " ] symbol."
      }
      return
    }
    index += 1
  }
  return
}
