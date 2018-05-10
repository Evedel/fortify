package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func ruleDeclaration(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
	indexInternal := 1
	resCode = UndefinedError
	errmsg = "ruleDeclaration: "
	resNode = TokenNodeVarDec()

	for indexInternal < len(ttail) {
		if ttail[indexInternal].Id == CarriageReturn {
			stopInd = indexInternal - 1
			resCode = Ok
			return
		} else if ttail[indexInternal].Id == Word {
			wordstr := ttail[indexInternal].Value
			if _, ok := Variables[wordstr]; !ok {
				Variables[wordstr] = Float
				resNode.List = append(resNode.List, TokenNodeVarId(wordstr))
			} else {
				resCode = AlreadyDeclared
				stopInd = indexInternal
				errmsg = "Variable [" + wordstr + "] has already been declared."
				return
			}
		} else if ttail[indexInternal].Id == Space {
			resNode.List = append(resNode.List, TokenNodeSpace())
		} else {
			resCode = UnexpectedArgument
			stopInd = indexInternal
			errmsg = "It is not allowed to use key word [ " + ttail[indexInternal].IdName + " ] as variable."
			return
		}
		indexInternal += 1
	}
	return
}
