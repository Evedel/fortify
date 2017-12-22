package dictionary

// import (
//   "say"
// )

func ruleCommentAll(ttail []Token) (correct bool, stopInd int, childs []TokenNode, errmsg string) {
  indexInternal := 0
  strval := ""
  for indexInternal < len(ttail) {
    if ttail[indexInternal].Id == CarriageReturn {
      correct = true
      stopInd = indexInternal
      errmsg = ""
      childs = append(
        childs,
        TokenNode{
          Token{String, "string", 0, 0, strval},
          nil})
      return
    } else {
      if ttail[indexInternal].Id == Word {
        strval += ttail[indexInternal].ValueStr
      } else {
        strval += ttail[indexInternal].IdName
      }
    }
    indexInternal += 1
  }
  return
}
