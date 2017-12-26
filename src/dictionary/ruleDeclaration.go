package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func ruleDeclaration(ttail []Token) (resCode int, stopInd int, childs []TokenNode, errmsg string) {
	indexInternal := 0
	resCode = Ok
	errmsg = ""

	for indexInternal < len(ttail) {
		if ttail[indexInternal].Id == CarriageReturn {
			childs = append(
				childs,
				TokenNode{
					Token{Expression, "expression", ""},
					append([]TokenNode{}, TokenNode{Token{CarriageReturn, "\\n", ""}, nil})})
			stopInd = indexInternal
			return
		} else if ttail[indexInternal].Id == Word {
			wordstr := ttail[indexInternal].Value
			if _, ok := Variables[wordstr]; !ok {
				Variables[wordstr] = Float
				childs = append(
					childs,
					TokenNode{
						Token{Expression, "expression", ""},
						append([]TokenNode{}, TokenNode{Token{VariableId, "VarId", wordstr}, nil})})
			} else {
				resCode = AlreadyDeclared
				stopInd = indexInternal
				errmsg = "Variable [" + wordstr + "] has already been declared."
				return
			}
		} else if ttail[indexInternal].Id == Space {
			childs = append(
				childs,
				TokenNode{
					Token{Expression, "expression", ""},
					append([]TokenNode{}, TokenNode{ttail[indexInternal], nil})})
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
