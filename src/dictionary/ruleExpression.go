package dictionary

// import(
// 	"say"
// )

import (
  "github.com/Evedel/fortify/src/say"
)

func RuleExpression(ttail []Token) (resCode int, stopInd int, resNode TokenNode, errmsg string) {
	thisName := "ruleExpression: "
	resCode = UndefinedError
	stopInd = 0
	index := 0
	chindex := 0

	for index < len(ttail) {
		say.L2("", ttail[index], "\n")
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
			// resCode, chindex, chchilds, errmsg = rulePrint(ttail[index+1:])
			stopInd = index + chindex + 1
			// resNode = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
			return
		} else if (tokenid == CommentTex) || (tokenid == CommentF90) {
			// resCode, chindex, chchilds, errmsg = ruleCommentF90Tex(ttail[index+1:])
			stopInd = index + chindex + 1
			// resNode = append([]TokenNode{}, TokenNode{ttail[index], chchilds})
			return
		} else if tokenid == CommentAll {
			resCode, chindex, resNode, errmsg = ruleCommentAll(ttail[index+1:])
			stopInd = index + chindex + 1
			return
		} else if tokenid == DeclarationVar {
			resCode, chindex, resNode, errmsg = ruleDeclaration(ttail[index+1:])
			stopInd = index + chindex + 1
			return
		} else if tokenid == Word {
			if _, ok := Variables[tokenvalstr]; ok {
				resCode, chindex, resNode, errmsg = ruleAssignment(ttail[index:])
				stopInd = index + chindex
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
