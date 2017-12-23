package dictionary

import("say")

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
