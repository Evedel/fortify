package syntaxer

import (
  "say"

  "dictionary"
)

func PrintTokenTree(TokenTree dictionary.TokenNode, level string) {
  say.L0(level + "> ", TokenTree.This, "\n")
  for ttch := range TokenTree.List {
    PrintTokenTree(TokenTree.List[ttch], level + "|--")
  }
}

func BuildTree(Tokenised []dictionary.Token) (TokenTree dictionary.TokenNode) {
  TokenTree.This = dictionary.Token{dictionary.Program, "program", 0, 0, ""}
  indx := 0
  for indx < len(Tokenised) {
    ok, sind, tch, errmsg := dictionary.RuleExpression(Tokenised[indx:])
    if ok {
      TokenTree.List = append(TokenTree.List,
        dictionary.TokenNode{
          dictionary.Token {
            dictionary.Expression, "expression", 0, 0, ""},
            tch})
      indx += sind + 1
    } else {
      say.L3(errmsg, "", "\n")
      indx  = len(Tokenised) + 1
    }
  }
  say.L1("TokenTree", "", "\n")
  PrintTokenTree(TokenTree, "--")
  return
}
