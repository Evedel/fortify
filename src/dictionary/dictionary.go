package dictionary

// import("say")

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
  "var": VarDeclaration}

var KeyWordBackslashReverse = map[int]string{}
var KeyWordBackslash  = map[string]int{
  "\\print": Print,
  "\\var": VarDeclaration }

var DataObjectReverse = map[int]string{}
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
