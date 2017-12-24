package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func ruleCommentAll(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
  resCode = Ok
  indexInternal := 0
  strval := ""
  errmsg = ""
  for indexInternal < len(ttail) {
    if ttail[indexInternal].Id == CarriageReturn {
      stopInd = indexInternal
      childs = append(
        childs,
        TokenNode{
          Token{String, "string", strval},
          nil})
      return
    } else {
      if ttail[indexInternal].Id == Word {
        strval += ttail[indexInternal].Value
      } else {
        strval += ttail[indexInternal].IdName
      }
    }
    indexInternal += 1
  }
  return
}
