package dictionary

// import(
//   "github.com/Evedel/fortify/src/say"
// )

func ruleOperand(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
	errmsg = "ruleOperand: "
	resCode = UndefinedError
	stopInd = 0
	index := 0
	chStopIndx := 0
	rhs := TokenNodeRightHS()

	for index < len(ttail) {
		tokenid := ttail[index].Id
		tokenvalstr := ttail[index].Value

		if tokenid == CarriageReturn {
			stopInd = index - 1
			resNode = rhs
			resCode = Ok
			return
		} else if (tokenid == Space) {
			rhs.List = append(rhs.List, TokenNodeSpace())
		} else if (tokenid == RoundBracketOpen) {
			iStop := -1
			iTmp := index
			for iTmp < len(ttail) {
				if (ttail[iTmp].Id == RoundBracketClose) {
					iStop = iTmp
					iTmp = len(ttail)
				}
				iTmp++
			}
			if (iStop == -1) {
				resCode = MissedRoundBracketClose
				errmsg += "Missed close round bracket"
				return
			}
			rhs.List = append(rhs.List, TokenNodeRoundBrackets())
			rhsInside := TokenNode{}
			resCode, chStopIndx, rhsInside, errmsg = ruleOperand(ttail[index+1:iStop])
			index += chStopIndx + 1
			rhs.List[len(rhs.List)-1].List = append(rhs.List[0].List, rhsInside)
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
		} else if tokenid == RoundBracketClose {
			resCode = MissedRoundBracketOpen
			errmsg += "Missed open round bracket"
			return
		} else {
			resCode = UnexpectedArgument
			errmsg += "Unexpected symbol in math expression: <|" + ttail[index].IdName + "|><|" + tokenvalstr + "|>"
			return
		}
		index += 1
	}

	stopInd = index
	resNode = rhs
	resCode = Ok
	return
}
