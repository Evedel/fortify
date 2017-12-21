package dictionary

// import (
//   "say"
// )

func RuleExpression(ttail []Token) (correct bool, stopInd int, childs []TokenNode, errmsg string) {
  correct = false
  stopInd = 0
  errmsg = "Not a correct expression."
  index := 0
  for index < len(ttail) {
    // say.L1("", ttail[index], "\n")
    if ttail[index].Id == Space {
      correct = true
      stopInd = index
      errmsg = ""
      childs = append([]TokenNode{}, TokenNode{ttail[index], nil})
      return
    } else if ttail[index].Id == CarriageReturn {
      correct = true
      stopInd = index
      errmsg = ""
      childs = append([]TokenNode{}, TokenNode{Token{ CarriageReturn, "\\n", 0, 0, ""}, nil})
      return
    } else if ttail[index].Id == Print {
      ok, is, chch, erms := RulePrint(ttail[index+1:])
      correct = ok
      if ok {
        stopInd = index + is + 1
        childs = append([]TokenNode{}, TokenNode{ttail[index], chch})
      } else {
        errmsg = erms
      }
      return
    } else if (ttail[index].Id == CommentTex) || (ttail[index].Id == CommentF90) {
      // say.L1("", ttail[index], "\n")
      // There should be valid expression between comment and \n
      CommentToken := TokenNode{ttail[index], []TokenNode{}}
      indexInternal := index + 1
      for indexInternal < len(ttail) {
        if ttail[indexInternal].Id == CarriageReturn {
          CommentToken.List = append(
            CommentToken.List,
            TokenNode{
              Token{ Expression, "expression", 0, 0, ""},
              append([]TokenNode{}, TokenNode{Token{ CarriageReturn, "\\n", 0, 0, ""}, nil})})
          childs = append(childs, CommentToken)
          correct = true
          stopInd = indexInternal
          errmsg = ""
          return
        } else {
          ok, is, chch, erms := RuleExpression(ttail[indexInternal:])
          if ok {
            CommentToken.List = append(CommentToken.List, TokenNode{Token{ Expression, "expression", 0, 0, ""}, chch})
            indexInternal += is
          } else {
            errmsg = erms
            return
          }
        }
        indexInternal += 1
      }
    } else if ttail[index].Id == CommentAll {
      CommentToken := TokenNode{ttail[index], []TokenNode{}}
      indexInternal := index + 1
      strval := ""
      for indexInternal < len(ttail) {
        if ttail[indexInternal].Id == CarriageReturn {
          correct = true
          stopInd = indexInternal
          errmsg = ""
          CommentToken.List = append(CommentToken.List, TokenNode{Token{ String, "string", 0, 0, strval}, nil})
          childs = append(childs, CommentToken)
          return
        } else {
          if ttail[indexInternal].Id == Word {
            strval += ttail[indexInternal].ValueStr
          } else {
            strval += ttail[indexInternal].IdName
          }
        }
        indexInternal += 1
      }
    // } else if ttail[index].IdName
    } else {
      if ttail[index].Id == Word {
        errmsg = "There is no defined rule for [ " + ttail[index].ValueStr + " ] symbol. " + errmsg
      } else {
        errmsg = "There is no defined rule for [ " + ttail[index].IdName + " ] symbol. " + errmsg
      }
      return
    }
    index += 1
  }
  return
}
