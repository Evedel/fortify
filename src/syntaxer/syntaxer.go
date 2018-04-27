package syntaxer

// import(
// 	"dictionary"
// 	// "say"
// )

import (
  // "github.com/Evedel/fortify/src/say"
  "github.com/Evedel/fortify/src/dictionary"
)

func BuildTree(Tokenised []dictionary.Token) (TokenTree dictionary.TokenNode, resCode int, errmsg string) {
	resCode = dictionary.UndefinedError
	TokenTree.This = dictionary.Token{dictionary.Program, "program", ""}
	indx := 0
	stopIndx := 0
	errmsg = ""
	expToken := dictionary.TokenNode{}

	for indx < len(Tokenised) {
		resCode, stopIndx, expToken, errmsg = dictionary.RuleExpression(Tokenised[indx:])
		if resCode == dictionary.Ok {
			TokenTree.List = append(TokenTree.List, expToken)
			indx += stopIndx + 1
		} else {
			indx = len(Tokenised) + 1
			return
		}
	}
	return
}
