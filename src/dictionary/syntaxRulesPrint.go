package dictionary

// import (
//   "say"
// )

func RulePrint(ttail []Token) (correct bool, stopInd int, childs TokenNode, errmsg string) {
  lentt := len(ttail)
  correct = false
  stopInd = 0
  errmsg = "Print default form is '\\print{a}'"
  childs.This = ttail[0]
  if (lentt < 3) {
    errmsg = "Not enough arguments. " + errmsg
    return
  } else {
    if ttail[1].Id != CurlyBracketOpen {
      errmsg = "There is no open bracket. " + errmsg
      return
    } else {
      childs.List = append(childs.List, TokenNode{ttail[1], nil})
      ind := 2
      for ind < lentt {
        if ttail[ind].Id == DoubleQuote {
          ok, sind, chchaild, erms := RuleString(ttail[ind:])
          if ok {
            ind += sind
            childs.List = append(childs.List, chchaild)
          } else {
            errmsg = erms
            return
          }
        // } else if other possible objects for print{
        // }
        } else if ttail[ind].Id == CurlyBracketClose {
          correct = true
          stopInd += ind
          childs.List = append(childs.List, TokenNode{ttail[ind], nil})
          errmsg = ""
          return
        } else {
          correct = false
          stopInd = ind
          errmsg = "Cannot use " + ttail[ind].IdName + " in print. " + errmsg
          return
        }
        ind += 1
      }
    }
  }
  errmsg = "There is no closing curly bracket in print. " + errmsg
  return
}

func RuleString(ttail []Token) (correct bool, stopInd int, childs TokenNode, errmsg string) {
  lentt := len(ttail)
  correct = false
  stopInd = 0
  errmsg = "Default string form is \"words words words\""
  if ttail[0].Id != DoubleQuote {
    errmsg = "There is no opening doublequote. " + errmsg
    return
  } else {
    ich := 1
    childs.List = append(childs.List, TokenNode{ttail[0], nil})
    for ich < lentt {
      childs.List = append(childs.List, TokenNode{ttail[ich], nil})
      if ttail[ich].Id == DoubleQuote {
        correct = true
        stopInd = ich
        errmsg = ""
        strval := ""
        for itch := 0; itch < len(childs.List); itch++ {
          if childs.List[itch].This.Id == Word {
            strval += childs.List[itch].This.ValueStr + " "
          } else {
            strval += childs.List[itch].This.IdName + " "
          }
        }
        childs.This = Token{ String, "string", 0, 0, strval}
        childs.List = []TokenNode{}
        return
      }
      ich += 1
    }
  }
  errmsg = "There is no closing doublequote. " + errmsg
  return
}
