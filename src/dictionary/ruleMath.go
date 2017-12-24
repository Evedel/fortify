package dictionary

// import (
//   "say"
// )

func ruleMath(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
  thisName := "ruleMath: "
  resCode = UndefinedError
  stopInd = 0
  index := 0
  inBrackets := false
  chStopIndx := 0
  // chchilds := []TokenNode{}
  // RootToken := TokenNode{}
  // say.L2("", ttail, "\n")
  for index < len(ttail) {
    // curtok := ttail[index]
    tokenid := ttail[index].Id
    tokenval := ttail[index].Value

    // say.L1("", ttail[index], "\n")
    if tokenid == Space {
      childs = append(childs, GetSpaceTokenNode())
    } else if tokenid == CarriageReturn {
      childs = append(childs, GetRNTokenNode())
      if (! inBrackets) {
        stopInd = index
        resCode = Ok
        return
      }
    } else if tokenid == Word {
      if _, ok := Variables[tokenval]; ok {
        resCode = TOTOTODO
        // resCode, chindex, chchilds, errmsg = ruleCalcExpression(ttail[index+1:])
        stopInd = index + chStopIndx + 1
        // childs = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
        return
      } else {
        chToken := TokenNode{}
        resCode, chStopIndx, chToken, errmsg = typeNumber(ttail[index:])
        if resCode == Ok {
          childs = append(childs, chToken)
          return
        } else {
          resCode = NotALanguageKeyWord
          errmsg = thisName + "There is no defined rule for [ " + tokenval + " ] symbol."
          return
        }
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
