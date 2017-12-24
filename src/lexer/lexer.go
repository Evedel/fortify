package lexer

import (
  // "say"
  "github.com/Evedel/fortify/src/dictionary"

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

  var t dictionary.Token
  t.Id = dictionary.None
  t.IdName = "none"
  word := ""
  for curpos < len(source) {
    word += string(source[curpos])

    if id, ok := dictionary.SpecialSymbol[word]; ok {
      t.Id = id
      t.IdName = word
    } else if id, ok := dictionary.KeyWordRaw[word]; ok {
      t.Id = id
      t.IdName = word
    } else if id, ok := dictionary.KeyWordBackslash[word]; ok {
      t.Id = id
      t.IdName = word
    }
    if t.Id == dictionary.None {
      for backpos := len(word)-1; backpos > -1; backpos-- {
        if id, ok := dictionary.SpecialSymbol[word[backpos:]]; ok {
          t.Id = id
          t.IdName = word[backpos:]
        } else if id, ok := dictionary.KeyWordRaw[word[backpos:]]; ok {
          t.Id = id
          t.IdName = word[backpos:]
        } else if id, ok := dictionary.KeyWordBackslash[word[backpos:]]; ok {
          t.Id = id
          t.IdName = word[backpos:]
        }
        if t.Id != dictionary.None {
          Tokenised = append(Tokenised,
            dictionary.Token{
              dictionary.Word,
              "word",
              word[:backpos]})
          break
        }
      }
    }

    if t.Id != dictionary.None {
      Tokenised = append(Tokenised, t)

      t.Id = dictionary.None
      t.IdName = "None"
      t.Value = ""

      word = ""
    }
    curpos += 1
  }
  return
}
