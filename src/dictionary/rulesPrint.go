package dictionary

// import (
//   "say"
// )

func RulePrint(ttail []Token) (correct bool, stopInd int, childs []TokenNode, errmsg string) {
  lentt := len(ttail)
  correct = false
  stopInd = 0
  errmsg = "Print default form is '\\print{var, \"str\"}'"
  if (lentt < 3) {
    errmsg = "Should contain at least one variable. " + errmsg
    return
  } else {
    if ttail[0].Id != CurlyBracketOpen {
      errmsg = "There is no open bracket. " + errmsg
      return
    } else {
      childs = append(childs, TokenNode{ttail[0], nil})
      stopInd = 1
      for stopInd < lentt {
        if ttail[stopInd].Id == DoubleQuote {
          ok, sind, chch, erms := RuleString(ttail[stopInd:])
          if ok {
            strval := ""
            for itch := 0; itch < len(chch); itch++ {
              if chch[itch].This.Id == Word {
                strval += chch[itch].This.ValueStr
              } else {
                strval += chch[itch].This.IdName
              }
            }
            stopInd += sind
            childs = append(childs,
              TokenNode{Token{ String, "string", 0, 0, strval},
              nil})
            // say.L0("Indx in print: ", stopInd, "\n")
          } else {
            errmsg = erms
            return
          }
        // } else if other possible objects for print{
        // }
        } else if ttail[stopInd].Id == CurlyBracketClose {
          correct = true
          childs = append(childs, TokenNode{ttail[stopInd], nil})
          errmsg = ""
          return
        } else {
          correct = false
          stopInd = stopInd
          errmsg = "Cannot use " + ttail[stopInd].IdName + " in print. " + errmsg
          return
        }
        stopInd += 1
      }
    }
  }
  errmsg = "There is no closing curly bracket in print. " + errmsg
  return
}

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
