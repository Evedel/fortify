package dictionary

import("say")

// all the possible words
const (
  Program    = iota
  Expression = iota

  None           = iota
  WrongSymbol    = iota

  CarriageReturn  = iota
  DoubleQuote     = iota
  Comma           = iota
  Space           = iota

  CommentTex = iota   // Not shown in tex, but compiled in f90
  CommentF90 = iota   // Not shown in f90, but compiled in tex
  CommentAll = iota   // Not compiled at all

  CurlyBracketOpen  = iota
  CurlyBracketClose = iota

  NumberInt = iota
  NumberFlt = iota
  Word      = iota
  String    = iota

  VariableId     = iota
  VariableInt    = iota
  VariableFloat  = iota
  VariableString = iota

  Print          = iota
  DeclarationVar = iota
)

var SpecialSymbolReverse = map[int]string{}
var SpecialSymbol = map[string]int{
  "\n": CarriageReturn,
  "%" : CommentTex,
  "!" : CommentF90,
  "#" : CommentAll,
  "\"": DoubleQuote,
  "{" : CurlyBracketOpen,
  "}" : CurlyBracketClose,
  "," : Comma,
  " " : Space }

var NeedbeMerroredReverse = map [int]string{
  CurlyBracketOpen  : "{",
  CurlyBracketClose : "}",
  CommentAll        : "#" }

var KeyWordRawReverse = map[int]string{}
var KeyWordRaw  = map[string]int{
  "print": Print,
  "var": DeclarationVar }

var KeyWordBackslashReverse = map[int]string{}
var KeyWordBackslash  = map[string]int{
  "\\print": Print,
  "\\var": DeclarationVar }

var DataObjectReverse = map[int]string{}
var DataObject = map[string]int{
  "word" : Word,
  "string" : String }

var Variables = map[string]int{}

type Token struct {
  Id int
  IdName string
  ValueInt int
  ValueFlt float64
  ValueStr string
}

type TokenNode struct {
  This Token
  List []TokenNode
}

func Init() {
  for key := range SpecialSymbol {
    SpecialSymbolReverse[SpecialSymbol[key]] = key
  }
  for key := range KeyWordRaw {
    KeyWordRawReverse[KeyWordRaw[key]] = key
  }
  for key := range DataObject {
    DataObjectReverse[DataObject[key]] = key
  }
  for key := range KeyWordBackslash {
    KeyWordBackslashReverse[KeyWordBackslash[key]] = key
  }
}

func PrintSyntaxTree(TokenTree TokenNode, level string) {
  say.L0(level + "> ", TokenTree.This, "\n")
  for ttch := range TokenTree.List {
    PrintSyntaxTree(TokenTree.List[ttch], level + "|--")
  }
}
