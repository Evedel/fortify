package dictionary

// import (
//   "say"
// )

func RuleExpression(ttail []Token) (correct bool, stopInd int, childs []TokenNode, errmsg string) {
  correct = false
  stopInd = 0
  errmsg = "Not a correct expression."
  index := 0
  chindex := 0
  chchilds := []TokenNode{}
  for index < len(ttail) {
    // say.L1("", ttail[index], "\n")
    if ttail[index].Id == Space {
      correct = true
      stopInd = index
      childs = append([]TokenNode{}, TokenNode{ttail[index], nil})
      return
    } else if ttail[index].Id == CarriageReturn {
      correct = true
      stopInd = index
      childs = append([]TokenNode{}, TokenNode{Token{ CarriageReturn, "\\n", 0, 0, ""}, nil})
      return
    } else if ttail[index].Id == Print {
      correct, chindex, chchilds, errmsg = RulePrint(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if (ttail[index].Id == CommentTex) || (ttail[index].Id == CommentF90) {
      correct, chindex, chchilds, errmsg = ruleCommentF90Tex(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if ttail[index].Id == CommentAll {
      correct, chindex, chchilds, errmsg = ruleCommentAll(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if ttail[index].Id == DeclarationVar {
      correct, chindex, chchilds, errmsg = ruleDeclaration(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else {
      if ttail[index].Id == Word {
        errmsg = "There is no defined rule for [ " + ttail[index].ValueStr + " ] symbol. " + errmsg
      } else {
        errmsg = "There is no defined rule for [ " + ttail[index].IdName + " ] symbol. " + errmsg
      }
      return
    }
    index += 1
  }
  return
}
