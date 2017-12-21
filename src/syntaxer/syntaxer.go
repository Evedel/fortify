package syntaxer

import (
  "say"

  "dictionary"
)

func PrintSyntaxTree(TokenTree dictionary.TokenNode, level string) {
  say.L0(level + "> ", TokenTree.This, "\n")
  for ttch := range TokenTree.List {
    PrintSyntaxTree(TokenTree.List[ttch], level + "|--")
  }
}

func BuildTree(Tokenised []dictionary.Token) (TokenTree dictionary.TokenNode, ok bool, errmsg string) {
  ok = true
  TokenTree.This = dictionary.Token{dictionary.Program, "program", 0, 0, ""}
  indx := 0
  errmsg = ""
  for indx < len(Tokenised) {
    okexp, sind, tch, cherrmsg := dictionary.RuleExpression(Tokenised[indx:])
    if okexp {
      TokenTree.List = append(TokenTree.List,
        dictionary.TokenNode{
          dictionary.Token {
            dictionary.Expression, "expression", 0, 0, ""},
            tch})
      indx += sind + 1
    } else {
      ok = false
      errmsg = cherrmsg
      indx  = len(Tokenised) + 1
      return
    }
  }
  return
}
