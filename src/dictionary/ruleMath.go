package dictionary

// import(
//   "say"
// )

// import(
//   "github.com/Evedel/fortify/src/say"
// )

func ruleMath(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
	thisName := "ruleMath: "
	resCode = UndefinedError
	stopInd = 0
	index := 0
	inBrackets := false
	chStopIndx := 0
	chchilds := []TokenNode{}
	// RootToken := TokenNode{}
	for index < len(ttail) {
		tokenthis := ttail[index]
		tokenid := ttail[index].Id
		tokenval := ttail[index].Value

		// say.L1("", ttail[index], "\n")
		if tokenid == Space {
			childs = append(childs, GetSpaceTokenNode())
		} else if tokenid == CarriageReturn {
			childs = append(childs, GetRNTokenNode())
			if !inBrackets {
				stopInd = index
				resCode = Ok
				return
			}
		} else if tokenid == Word {
			if _, ok := Variables[tokenval]; ok {
				resCode = Ok
				stopInd = index + chStopIndx + 1
				childs = append(childs, ToVarIdTokenNode(tokenthis))
			} else {
				chToken := TokenNode{}
				resCode, chStopIndx, chToken, errmsg = typeNumber(ttail[index:])
				if resCode == Ok {
					stopInd = index + chStopIndx
					childs = append(childs, chToken)
					return
				} else {
					resCode = NotALanguageKeyWord
					errmsg = thisName + "There is no defined rule for [ " + tokenval + " ] symbol."
					return
				}
			}
		} else if (tokenid == Addition) ||
			(tokenid == Substraction) ||
			(tokenid == Multiplication) ||
			(tokenid == Division) {
			newChilds := []TokenNode{TokenNode{tokenthis,
				[]TokenNode{ExpressionTokenNode(childs)}}}
			resCode, chStopIndx, chchilds, errmsg = ruleMath(ttail[index+1:])
			if resCode == Ok {
				stopInd = index + chStopIndx + 1
				newChilds[0].List = append(newChilds[0].List, ExpressionTokenNode(chchilds))
				childs = newChilds
				return
			} else {
				return
			}
		} else {
			resCode = NotALanguageKeyWord
			errmsg = thisName + "There is no defined rule for [ " + ttail[index].IdName + " ] symbol."
			return
		}
		index += 1
	}
	return
}
