package compile

import(
  "dictionary"
)
const (
  ltfort = 0
  ltclang = 1
)

var CompiledTokens = map[int][]string{
  dictionary.Print: []string{"", ""} }
