package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func typeString(ttail []Token) (resCode int, stopInd int, resToken TokenNode, errmsg string) {
  lentt := len(ttail)
  resCode = UndefinedError
  stopInd = 0
  errmsg = ""
  childs := []TokenNode{}
  // errmsg = "Default string form is \"word word word ...\""
  if ttail[0].Id != DoubleQuote {
    resCode = LostBracket
    errmsg = "There is no opening doublequote. " + errmsg
    return
  } else {
    ich := 1
    childs = append(childs, TokenNode{ttail[0], nil})
    for ich < lentt {
      childs = append(childs, TokenNode{ttail[ich], nil})
      if ttail[ich].Id == DoubleQuote {
        resCode = Ok
        stopInd = ich
        strval := ""
        for itch := 0; itch < len(childs); itch++ {
          if childs[itch].This.Id == Word {
            strval += childs[itch].This.Value
          } else {
            strval += childs[itch].This.IdName
          }
        }
        resToken = TokenNode{Token{ String, "string", strval}, nil}
        return
      }
      ich += 1
    }
  }
  resCode = LostBracket
  errmsg = "There is no closing doublequote in string. Default string is \"word word ...\"."
  return
}
