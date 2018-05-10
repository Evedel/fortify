package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func RuleExpression(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
	thisName := "ruleExpression: "
	resCode = UndefinedError
	stopInd = 0
	index := 0
	chindex := 0

	for index < len(ttail) {
		// say.L0("Rule Expression: ", ttail[index], "\n")
		tokenid := ttail[index].Id
		tokenvalstr := ttail[index].Value
		if tokenid == Space {
			resCode = Ok
			stopInd = index
			resNode = TokenNodeSpace()
			return
		} else if tokenid == CarriageReturn {
			resCode = Ok
			stopInd = index
			resNode = TokenNodeReturn()
			return
		} else if tokenid == Print {
			resCode, chindex, resNode, errmsg = rulePrint(ttail[index:])
			stopInd = index + chindex
			return
		} else if (tokenid == DontCompileTex) || (tokenid == DontCompileF90) {
			resCode, chindex, resNode, errmsg = ruleDontCompileF90Tex(ttail[index:])
			stopInd = index + chindex
			return
		} else if tokenid == CommentAll {
			resCode, chindex, resNode, errmsg = ruleCommentAll(ttail[index:])
			stopInd = index + chindex
			return
		} else if tokenid == DeclarationVar {
			resCode, chindex, resNode, errmsg = ruleDeclaration(ttail[index:])
			stopInd = index + chindex
			return
		} else if tokenid == Word {
			if _, ok := Variables[tokenvalstr]; ok {
				resCode, chindex, resNode, errmsg = ruleAssignment(ttail[index:])
				stopInd = index + chindex
				// say.L0("Rule Assignment stopIndex: ", ttail[stopInd], "\n")
				return
			} else {
				resCode = NotALanguageKeyWord
				errmsg = thisName + "There is no defined rule for [ " + tokenvalstr + " ] symbol."
				return
			}
		} else {
			resCode = NotALanguageKeyWord
			errmsg = thisName + "There is no defined rule for [ " + ttail[index].IdName + " ] symbol."
			return
		}
	}
	return
}
