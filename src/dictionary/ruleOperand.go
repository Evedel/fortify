package dictionary

// import(
//   "github.com/Evedel/fortify/src/say"
// )

func ruleOperand(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
	thisName := "ruleOperand: "
	resCode = UndefinedError
	stopInd = 0
	index := 0
	chStopIndx := 0
	rhs := TokenNode{}
  op := TokenNodeOperand()

	for index < len(ttail) {
		tokenid := ttail[index].Id
		tokenvalstr := ttail[index].Value

		if tokenid == CarriageReturn {
			stopInd = index
			resCode = Ok
			resNode = op
			return
		} else if (tokenid == Addition) ||
		(tokenid == Substraction) ||
		(tokenid == Multiplication) ||
		(tokenid == Division) {
			resNode = TokenNodeFromToken(ttail[index])
			resNode.List = append(resNode.List, op)
			resCode, chStopIndx, rhs, errmsg = ruleOperand(ttail[index+1:])
			stopInd = index + chStopIndx
			resNode.List = append(resNode.List, rhs)
			return
		}
		if tokenid == Word {
      if _, ok := Variables[tokenvalstr]; ok {
				op.List = append(op.List, TokenNodeVarId(tokenvalstr))
      } else {
        resCode = NotALanguageKeyWord
        errmsg = thisName + "There is no defined rule for [ " + tokenvalstr + " ] symbol."
        return
      }
		} else {
			op.List = append(op.List, TokenNodeFromToken(ttail[index]))
		}

		index += 1
	}
	thisName += errmsg
	errmsg = thisName
	return
}
