package compile

import(
  "github.com/Evedel/fortify/src/dictionary"
)
const (
  ltfort = 0
  ltclang = 1
)

var CompiledTokens = map[int][]string{
  dictionary.Print: []string{"", ""} }
