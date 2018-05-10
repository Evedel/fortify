package compile

import(
	"github.com/Evedel/fortify/src/dictionary"
	// "github.com/Evedel/fortify/src/say"
)

func reduceToFortran(HeadToken dictionary.TokenNode) (newHeadToken dictionary.TokenNode, reduse bool) {
  newHeadToken = dictionary.TokenNode{}
  reduse = true

  if (HeadToken.This.Id == dictionary.Space) ||
      (HeadToken.This.Id == dictionary.CarriageReturn) ||
      (HeadToken.This.Id == dictionary.CommentAll) ||
			(HeadToken.This.Id == dictionary.DontCompileF90) ||
      (HeadToken.This.Id == dictionary.DeclarationVar) {
        return
	} else {
    childs := []dictionary.TokenNode{}
    if len(HeadToken.List) == 0 {
			if (HeadToken.This.Id == dictionary.LeftHS) ||
				(HeadToken.This.Id == dictionary.RightHS) {
					return
			} else {
				newHeadToken = dictionary.TokenNode{HeadToken.This, nil}
	      reduse = false
	      return				
			}
    } else {
      for _, subtoken := range HeadToken.List {
        reducedSubToken, needReduce := reduceToFortran(subtoken)
        if (! needReduce) {
          childs = append(childs, reducedSubToken)
        }
      }
      if (len(childs) == 0) && (HeadToken.This.Id != dictionary.Program) {
        return
      } else {
        newHeadToken = dictionary.TokenNode{HeadToken.This, childs}
        reduse = false
        return
      }
    }
  }
}
