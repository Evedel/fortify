package dictionary

import(
	"github.com/Evedel/fortify/src/say"
)

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
	say.L0(level+"> ", TokenTree.This, "\n")
	for ttch := range TokenTree.List {
		PrintSyntaxTree(TokenTree.List[ttch], level+"+   ")
	}
}

func ToVarIdTokenNode(t Token) TokenNode {
	return TokenNode{Token{VariableId, "VarId", t.Value}, nil}
}

func roundBracketOpenTN() TokenNode {
	return TokenNode{Token{RoundBracketOpen, "(", "("}, nil}
}

func roundBracketCloseTN() TokenNode {
	return TokenNode{Token{RoundBracketOpen, ")", ")"}, nil}
}
