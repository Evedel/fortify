package dictionary

// import (
//   "time"
//   "github.com/Evedel/fortify/src/say"
// )

func ruleDontCompileF90Tex(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
  resCode = UndefinedError
  indexChild := 0
  subNode := TokenNode{}
  errmsg += "ruleDontCompileF90Tex: "

  resNode = TokenNodeFromToken(ttail[0])
  indexInternal := 1
  for indexInternal < len(ttail) {
  	if ttail[indexInternal].Id == CarriageReturn {
  		resCode = Ok
  		stopInd = indexInternal - 1
  		return
  	} else {
  		resCode, indexChild, subNode, errmsg = RuleExpression(ttail[indexInternal:])
  		if resCode == Ok {
  			resNode.List = append(resNode.List, subNode)
  			indexInternal += indexChild
        indexInternal++
  		} else {
  			return
  		}
  	}
  }
  return
}
