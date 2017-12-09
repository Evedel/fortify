package dictionary

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
  CommentAll = iota   // Not compiled ta all

  CurlyBracketOpen  = iota
  CurlyBracketClose = iota

  NumberInt = iota
  NumberFlt = iota
  Word      = iota
  String    = iota

  VariableId     = iota
  Print          = iota
  VarDeclaration = iota
)

// Symbols
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

var KeyWordRaw  = map[string]int{
  "print": Print,
  "var": VarDeclaration}

var KeyWordBackslash  = map[string]int{
  "\\print": Print,
  "\\var": VarDeclaration }

var DataObject = map[string]int{
  "word" : Word,
  "string" : String }

type Token struct {
  Id int
  IdName string
  ValueInt int
  ValueFlt float64
  ValueStr string
}

type TokenNode struct {
  This Token
  // Parent TokenNode
  List []TokenNode
}

// ------============++++++++++++++============------ //
type RuleInterface func ([]Token)
