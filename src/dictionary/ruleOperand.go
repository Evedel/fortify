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
	rhs := TokenNodeRightHS()

	errmsg += thisName
	for index < len(ttail) {
		tokenid := ttail[index].Id
		tokenvalstr := ttail[index].Value

		if tokenid == CarriageReturn {
			stopInd = index
			resNode = rhs
			resCode = Ok
			return
		} else if (tokenid == Space) {
			rhs.List = append(rhs.List, TokenNodeSpace())
		} else if (tokenid == RoundBracketOpen) {
			iStop := index
			iTmp := index
			for iTmp < len(ttail) {
				if (ttail[iTmp].Id == RoundBracketClose) {
					iStop = iTmp
					iTmp = len(ttail)
				}
				iTmp++
			}
			resNode = TokenNodeRightHS()
			resNode.List = append([]TokenNode{}, TokenNodeFromToken(ttail[index]))
			resCode, chStopIndx, rhs, errmsg = ruleOperand(ttail[index+1:iStop])
			stopInd = index + chStopIndx + 1
			resNode.List = append(resNode.List, rhs)
			resNode.List = append(resNode.List, TokenNodeFromToken(ttail[iStop]))
			return
		} else if (tokenid == Addition) ||
		(tokenid == Substraction) ||
		(tokenid == Multiplication) ||
		(tokenid == Division) {
			resNode = TokenNodeRightHS()
			operator := TokenNodeFromToken(ttail[index])
			lhs := TokenNodeRHS2LHS(rhs)
			operator.List = append([]TokenNode{}, lhs)
			resCode, chStopIndx, rhs, errmsg = ruleOperand(ttail[index+1:])
			stopInd = index + chStopIndx + 1
			operator.List = append(operator.List, rhs)
			resNode.List = append(resNode.List, operator)
			return
		} else if tokenid == Word {
      if _, ok := Variables[tokenvalstr]; ok {
				rhs.List = append(rhs.List, TokenNodeVarId(tokenvalstr))
      } else {
				rhs.List = append(rhs.List, typeStatic(ttail[index]))
      }
		} else {
			resCode = UnexpectedArgument
			errmsg = "Unexpected symbol in math expression: <|" + ttail[index].IdName + "|><|" + tokenvalstr + "|>"
			return
		}
		index += 1
	}

	stopInd = index
	resNode = rhs
	resCode = Ok
	return
}
