package syntaxer

import(
  // "say"

  "dictionary"
)

func BuildTree(Tokenised []dictionary.Token) (TokenTree dictionary.TokenNode, resCode int, errmsg string) {
  resCode = dictionary.UndefinedError
  TokenTree.This = dictionary.Token{dictionary.Program, "program", ""}
  indx := 0
  stopIndx := 0
  errmsg = ""
  chchilds := []dictionary.TokenNode{}

  for indx < len(Tokenised) {
    resCode, stopIndx, chchilds, errmsg = dictionary.RuleExpression(Tokenised[indx:])
    if resCode == dictionary.Ok {
      TokenTree.List = append(TokenTree.List,
        dictionary.TokenNode{
          dictionary.Token {
            dictionary.Expression, "expression", ""},
            chchilds})
      indx += stopIndx + 1
    } else {
      indx  = len(Tokenised) + 1
      return
    }
  }
  return
}
