package lexer

import (
  // "say"
  "dictionary"

  // "unicode"
)

// constants for token guessing
const (
  s_udf = iota
  s_str = iota
  s_nmb = iota
  s_cmd = iota
)

func Tokenise(source string) (Tokenised []dictionary.Token) {
  curpos := 0
  var word string

  var t dictionary.Token
  t.Id = dictionary.None
  t.IdName = "none"

  for curpos < len(source) {
    char := string(source[curpos])

    if (char == " ") {
      if len(word) > 0 {
        if id, ok := dictionary.KeyWordRaw[word]; ok {
          t.Id = id
          t.IdName = word
        } else if id, ok := dictionary.KeyWordBackslash[word]; ok {
          t.Id = id
          t.IdName = word
        } else if t.Id == dictionary.None {
          t.Id       = dictionary.Word
          t.IdName   = "word"
          t.ValueStr = word
        }
      } else {
        char = ""
      }
    } else {
      if id, ok := dictionary.SpecialSymbol[char]; ok {
        t.Id = id
        t.IdName = char
        // word is not empty but there is special symbol
        if len(word) != 0 {
          var t2 dictionary.Token
          t2.Id       = dictionary.Word
          t2.IdName   = "word"
          t2.ValueStr = word
          // say.L1("|" + t2.IdName + "|", t2.Id, "|" + t2.ValueStr + "|\n" )
          Tokenised = append(Tokenised, t2)
        }
      }
    }

    if t.Id != dictionary.None {
      // say.L1("|" + t.IdName + "|", t.Id, "|" + t.ValueStr + "|\n" )
      Tokenised = append(Tokenised, t)

      t.Id = dictionary.None
      t.IdName = "None"
      t.ValueInt = 0
      t.ValueFlt = 0.0
      t.ValueStr = ""

      word = ""
      char = ""
    } else {
      word += char
    }
    curpos += 1
  }
  return
}
