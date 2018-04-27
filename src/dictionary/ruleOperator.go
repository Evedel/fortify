package dictionary

// import(
//   "github.com/Evedel/fortify/src/say"
// )

func ruleOperator(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
	thisName := "ruleAssignment: "
	resCode = UndefinedError
	stopInd = 0
	index := 1
	chStopIndx := 0
	lhs := []TokenNode{}
	rhs := TokenNode{}
	lhs = append(lhs, TokenNodeVarId(ttail[0].Value))

	for index < len(ttail) {
		tokenthis := ttail[index]
		tokenid := ttail[index].Id

		if tokenid == Space {
			lhs = append(lhs, TokenNodeSpace())
		} else if tokenid == Assignment {
			resNode = TokenNodeAssignment()
			resNode.List = append(resNode.List, TokenNodeOperand())
			// resCode, chStopIndx, rhs, errmsg = ruleMathExpression(ttail[index+1:])
			stopInd = index + chStopIndx + 1
			resNode.List = append(resNode.List, rhs)
			return
		} else if tokenid == CarriageReturn {
			resCode = NotEnoughArguments
			errmsg = thisName + "There is not enough arguments in assignment"
			return
		} else {
			resCode = NotALanguageKeyWord
			if tokenid == Word {
				errmsg = thisName + "There is no defined rule for [ " + tokenthis.Value + " ] symbol."
			} else {
				errmsg = thisName + "There is no defined rule for [ " + tokenthis.IdName + " ] symbol."
			}
			return
		}
		index += 1
	}
	return
}
