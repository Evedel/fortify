package dictionary

// import (
//   "say"
// )

func ruleDeclaration(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
  indexInternal := 0
  resCode = Ok
  errmsg = ""

  for indexInternal < len(ttail) {
    if ttail[indexInternal].Id == CarriageReturn {
      childs = append(
        childs,
        TokenNode{
          Token{ Expression, "expression", 0, 0, ""},
          append([]TokenNode{}, TokenNode{Token{ CarriageReturn, "\\n", 0, 0, ""}, nil})})
      stopInd = indexInternal
      return
    } else if ttail[indexInternal].Id == Word {
      wordstr := ttail[indexInternal].ValueStr
      if _, ok := Variables[wordstr]; !ok {
        Variables[wordstr] = VariableFloat
        childs = append(
          childs,
          TokenNode{
            Token{ Expression, "expression", 0, 0, ""},
            append([]TokenNode{}, TokenNode{Token{ VariableId, "VarId", VariableFloat, 0, wordstr}, nil})})
      } else {
        resCode = AlreadyDeclared
        stopInd = indexInternal
        errmsg = "Variable [" + wordstr + "] has already been declared."
        return
      }
    } else if ttail[indexInternal].Id == Space {
      childs = append(
        childs,
        TokenNode{
          Token{ Expression, "expression", 0, 0, ""},
          append([]TokenNode{}, TokenNode{ttail[indexInternal], nil})})
    } else {
      resCode = UnexpectableArgument
      stopInd = indexInternal
      errmsg = "It is not allowed to use key word [ " + ttail[indexInternal].IdName + " ] as variable."
      return
    }
    indexInternal += 1
  }
  return
}
