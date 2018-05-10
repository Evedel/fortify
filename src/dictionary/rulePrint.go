package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func rulePrint(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {

	indexInternal := 1
	resCode = UndefinedError
	errmsg = "rulePrint: "
  resNode = TokenNodePrint()

  tokindx := ttail[indexInternal].Id
  indexInternal++
  if tokindx != CurlyBracketOpen {
    resCode = LostBracket
    errmsg += "There is no open curly bracket."
    return
  } else {
    resNode.List = append(resNode.List, TokenNodeFromToken(ttail[1]))
    for indexInternal < len(ttail) {
      tokindx = ttail[indexInternal].Id
  		if tokindx == DoubleQuote {
  			chStopIndx := 0
  			chtoken := TokenNode{}
  			cherrmsg := ""
  			resCode, chStopIndx, chtoken, cherrmsg = typeString(ttail[indexInternal:])
  			if resCode == Ok {
  				indexInternal += chStopIndx
  				resNode.List = append(resNode.List, chtoken)
  			} else {
  				errmsg = cherrmsg
  				return
  			}
  		} else if tokindx == Space {
  			resNode.List = append(resNode.List, TokenNodeSpace())
  		} else if tokindx == CurlyBracketClose {
  			resCode = Ok
  			resNode.List = append(resNode.List, TokenNodeFromToken(ttail[indexInternal]))
        stopInd = indexInternal
  			return
  		} else if tokindx == CarriageReturn {
  			// do nothing
  		} else if tokindx == Word {
  			if _, ok := Variables[ttail[indexInternal].Value]; ok {
  				resNode.List = append(resNode.List, ToVarIdTokenNode(ttail[indexInternal]))
  			} else {
  				resCode = UnexpectedArgument
  				errmsg = "Undefined token [ " + ttail[indexInternal].IdName + " ]. " + errmsg
  				return
  			}
  		} else {
  			resCode = UnexpectedArgument
  			errmsg = "Undefined token [ " + ttail[indexInternal].IdName + " ]. " + errmsg
  			return
  		}
      indexInternal++
    }
  }

	resCode = LostBracket
	errmsg = "There is no closing curly bracket. " + errmsg
	return
}
