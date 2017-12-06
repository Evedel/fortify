package syntaxer

import (
  "say"

  "dictionary"
)

func PrintTokenTree(TokenTree dictionary.TokenNode, level string) {
  say.L0(level, TokenTree.This, "\n")
  for ttch := range TokenTree.List {
    PrintTokenTree(TokenTree.List[ttch], level + "-->")
  }
}

func BuildTree(Tokenised []dictionary.Token) (TokenTree dictionary.TokenNode) {

  // say.L0("", Tokenised[1:], "\n")
  // RulePrint(ttail []Token) (correct bool, stopInd int, childs TokenNode, errmsg string) {
  ok, sind, TokenTree, errmsg := dictionary.RulePrint(Tokenised[1:])
  say.L1("", ok, "\n")
  say.L1("", sind, "\n")
  say.L1("", errmsg, "\n")
  say.L1("TokenTree", "", "\n")
  PrintTokenTree(TokenTree, "-->")
  return
}
