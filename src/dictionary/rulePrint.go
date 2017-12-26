package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func rulePrint(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
	lentt := len(ttail)
	resCode = Ok
	stopInd = 0
	errmsg = "Supposed format print{\"string\" decleradVarable ...}."
	if lentt < 3 {
		errmsg = "Should contain at least one variable. " + errmsg
		return
	} else {
		if ttail[0].Id != CurlyBracketOpen {
			resCode = LostBracket
			errmsg = "There is no open curly bracket. " + errmsg
			return
		} else {
			childs = append(childs, TokenNode{ttail[0], nil})
			stopInd = 1
			for stopInd < lentt {
				if ttail[stopInd].Id == DoubleQuote {
					chStopIndx := 0
					chtoken := TokenNode{}
					cherrmsg := ""
					resCode, chStopIndx, chtoken, cherrmsg = typeString(ttail[stopInd:])
					if resCode == Ok {
						stopInd += chStopIndx
						childs = append(childs, chtoken)
					} else {
						errmsg = cherrmsg
						return
					}
				} else if ttail[stopInd].Id == Space {
					childs = append(childs, GetSpaceTokenNode())
				} else if ttail[stopInd].Id == CurlyBracketClose {
					resCode = Ok
					childs = append(childs, TokenNode{ttail[stopInd], nil})
					// PrintSyntaxTree(TokenNode{Token{}, childs}, "")
					return
				} else if ttail[stopInd].Id == CarriageReturn {
					// do nothing
				} else if ttail[stopInd].Id == Word {
					if _, ok := Variables[ttail[stopInd].Value]; ok {
						childs = append(childs, ToVarIdTokenNode(ttail[stopInd]))
					} else {
						resCode = UnexpectedArgument
						errmsg = "Undefined token [ " + ttail[stopInd].IdName + " ]. " + errmsg
						return
					}
				} else {
					resCode = UnexpectedArgument
					errmsg = "Undefined token [ " + ttail[stopInd].IdName + " ]. " + errmsg
					return
				}
				stopInd += 1
			}
		}
	}
	resCode = LostBracket
	errmsg = "There is no closing curly bracket. " + errmsg
	return
}
