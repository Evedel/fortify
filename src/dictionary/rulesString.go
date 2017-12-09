package dictionary

// import (
//   "say"
// )

func RuleString(ttail []Token) (correct bool, stopInd int, childs []TokenNode, errmsg string) {
  lentt := len(ttail)
  correct = false
  stopInd = 0
  errmsg = "Default string form is \"words words words\""
  if ttail[0].Id != DoubleQuote {
    errmsg = "There is no opening doublequote. " + errmsg
    return
  } else {
    ich := 1
    childs = append(childs, TokenNode{ttail[0], nil})
    for ich < lentt {
      childs = append(childs, TokenNode{ttail[ich], nil})
      if ttail[ich].Id == DoubleQuote {
        correct = true
        stopInd = ich
        errmsg = ""
        return
      }
      ich += 1
    }
  }
  errmsg = "There is no closing doublequote. " + errmsg
  return
}
