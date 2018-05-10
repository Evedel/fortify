package dictionary

// import (
//   "github.com/Evedel/fortify/src/say"
// )

func ruleCommentAll(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
	resCode = Ok
	indexInternal := 1
	strval := ""
	errmsg = ""
	resNode = TokenNodeCommentAll()

	for indexInternal < len(ttail) {
		if ttail[indexInternal].Id == CarriageReturn {
			stopInd = indexInternal - 1
			resNode.List = append(resNode.List, TokenNodeString(strval))
			return
		} else {
			if ttail[indexInternal].Id == Word {
				strval += ttail[indexInternal].Value
			} else {
				strval += ttail[indexInternal].IdName
			}
		}
		indexInternal += 1
	}
	return
}
