package dictionary

// import(
//   "github.com/Evedel/fortify/src/say"
// )

func ruleAssignment(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
	thisName := "ruleAssignment: "
	resCode = UndefinedError
	stopInd = 0
	index := 1
	chStopIndx := 0
	lhs := []TokenNode{}
	rhs := TokenNode{}
	lhs = append(lhs, TokenNodeVarId(ttail[0].Value))

	for index < len(ttail) {
		tokenid := ttail[index].Id
		tokenvalstr := ttail[index].Value

		if tokenid == Space {
			lhs = append(lhs, TokenNodeSpace())
		} else if tokenid == Assignment {
			resNode = TokenNodeAssignment()
			lhsEnd := TokenNodeOperand()
			lhsEnd.List = append(lhs)
			resNode.List = append(resNode.List, lhsEnd)
			resCode, chStopIndx, rhs, errmsg = ruleOperand(ttail[index+1:])
			stopInd = index + chStopIndx + 1
			if len(rhs.List) != 0 {
				rhsEnd := TokenNodeOperand()
				rhsEnd.List = append(rhsEnd.List, rhs)
				resNode.List = append(resNode.List, rhsEnd)
			} else {
				resNode.List = append(resNode.List, rhs)
			}
			return
		} else if tokenid == CarriageReturn {
			resCode = NotEnoughArguments
			errmsg = thisName + "There is not enough arguments in assignment"
			return
		} else if tokenid == Word {
      if _, ok := Variables[tokenvalstr]; ok {
				resCode = UnexpectedArgument
				errmsg = thisName + "Cannot use two variables on the left of assignment"
				return
      } else {
        resCode = NotALanguageKeyWord
        errmsg = thisName + "There is no defined rule for [ " + tokenvalstr + " ] symbol."
        return
      }
		} else {
			resCode = NotALanguageKeyWord
			errmsg = thisName + "There is no defined rule for [ " + tokenvalstr + " ] symbol."
			return
		}
		index += 1
	}
	return
}
