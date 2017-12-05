package dictionary

// all the possible words
const (
  None           = iota
  WrongSymbol    = iota
  CarriageReturn = iota
  CommentTex     = iota
  CommentF90     = iota
  DoubleQuote    = iota
  NumberInt      = iota
  numberFlt      = iota
  VariableId     = iota
  Word           = iota
  Print          = iota
  VarDeclaration = iota
)

// Symbols
var SpecialSymbol = map[string]int{
  "\n": CarriageReturn,
  "!" : CommentTex,
  "#" : CommentF90,
  "\"": DoubleQuote }

// Key words on fortify
var KeyWordRaw  = map[string]int{
  "print": Print,
  "var": VarDeclaration}

// Key words of fortify with backslash
var KeyWordBackslash  = map[string]int{
  "\\print": Print,
  "\\var": VarDeclaration }

var DataObject = map[string]int{
  "string" : Word }

type Token struct {
    Id int
    IdName string
    ValueInt int
    ValueFlt float64
    ValueStr string }

func NameFromId(Id int) (key string, ok bool) {
  for k, v := range SpecialSymbol {
    if v == Id {
      key = k
      ok = true
      return
    }
  }
  for k, v := range KeyWordRaw {
    if v == Id {
      key = k
      ok = true
      return
    }
  }
  for k, v := range KeyWordBackslash {
    if v == Id {
      key = k
      ok = true
      return
    }
  }
  for k, v := range DataObject {
    if v == Id {
      key = k
      ok = true
      return
    }
  }
  return
}
