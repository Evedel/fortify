package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func ruleAssignment(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
  thisName := "ruleAssignment: "
  resCode = UndefinedError
  stopInd = 0
  index := 1
  chStopIndx := 0
  chchilds := []TokenNode{}
  rootToken := TokenNode{}
  lhs := GetExpressionTokenNode()
  lhs.List = append(lhs.List, ToVarIdTokenNode(ttail[0]))

  for index < len(ttail) {
    curtok := ttail[index]
    tokenid := ttail[index].Id

    if tokenid == Space {
      lhs.List = append(lhs.List, GetSpaceTokenNode())
    } else if tokenid == CarriageReturn {
      lhs.List = append(lhs.List, GetRNTokenNode())
    } else if tokenid == Assignment {
      rootToken = TokenNode{curtok, []TokenNode{lhs}}
      resCode, chStopIndx, chchilds, errmsg = ruleMath(ttail[index+1:])
      stopInd = index + chStopIndx + 1
      rootToken.List = append(rootToken.List,
        TokenNode{
          Token{ Expression, "expression", ""},
          chchilds})
      childs = append([]TokenNode{}, rootToken)
      return
    } else {
      resCode = NotALanguageKeyWord
      errmsg = thisName + "There is no defined rule for [ " + ttail[index].IdName + " ] symbol."
      return
    }
    index += 1
  }
  return
}
