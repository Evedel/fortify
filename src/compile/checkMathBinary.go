package compile

import(
	"github.com/Evedel/fortify/src/dictionary"
	// "github.com/Evedel/fortify/src/say"
)

func checkMathBinary(HeadToken dictionary.TokenNode) (resCode int) {
  resCode = dictionary.UndefinedError
  dictionary.PrintSyntaxTree(HeadToken, "")
  // say.L1("", HeadToken.List, "\n")
  // for _, subtoken := range HeadToken.List {
  //
  // }
  return
}
