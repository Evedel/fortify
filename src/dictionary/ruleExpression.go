package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func RuleExpression(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
  thisName := "ruleExpression: "
  resCode = UndefinedError
  stopInd = 0
  index := 0
  chindex := 0
  chchilds := []TokenNode{}
  for index < len(ttail) {
    tokenid := ttail[index].Id
    tokenvalstr := ttail[index].Value
    if tokenid == Space {
      resCode = Ok
      stopInd = index
      childs = append([]TokenNode{}, TokenNode{ttail[index], nil})
      return
    } else if tokenid == CarriageReturn {
      resCode = Ok
      stopInd = index
      childs = append([]TokenNode{}, TokenNode{Token{ CarriageReturn, "\\n", ""}, nil})
      return
    } else if tokenid == Print {
      resCode, chindex, chchilds, errmsg = RulePrint(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if (tokenid == CommentTex) || (tokenid == CommentF90) {
      resCode, chindex, chchilds, errmsg = ruleCommentF90Tex(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if tokenid == CommentAll {
      resCode, chindex, chchilds, errmsg = ruleCommentAll(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if tokenid == DeclarationVar {
      resCode, chindex, chchilds, errmsg = ruleDeclaration(ttail[index+1:])
      stopInd = index + chindex + 1
      childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
      return
    } else if tokenid == Word {
      if _, ok := Variables[tokenvalstr]; ok {
        resCode, chindex, chchilds, errmsg = ruleAssignment(ttail[index:])
        stopInd = index + chindex + 1
        childs = chchilds
        return
      } else {
        resCode = NotALanguageKeyWord
        errmsg = thisName + "There is no defined rule for [ " + tokenvalstr + " ] symbol."
        return
      }
    } else {
      resCode = NotALanguageKeyWord
      errmsg = thisName + "There is no defined rule for [ " + ttail[index].IdName + " ] symbol."
      return
    }
    index += 1
  }
  return
}
