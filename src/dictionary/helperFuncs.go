package dictionary

import(
	"github.com/Evedel/fortify/src/say"
)
// import(
// 	"say"
// )

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
		PrintSyntaxTree(TokenTree.List[ttch], level+"|--")
	}
}

func GetRNTokenNode() TokenNode {
	return TokenNode{Token{CarriageReturn, "\\n", ""}, nil}
}

func ExpressionTokenNode(ch []TokenNode) TokenNode {
	return TokenNode{
		Token{Expression, "expression", ""}, ch}
}

func ExpressionFromToken(t Token, ch []TokenNode) TokenNode {
	return TokenNode{
		Token{Expression, "expression", ""},
		[]TokenNode{TokenNode{t, ch}}}
}

func GetSpaceTokenNode() TokenNode {
	return TokenNode{Token{Space, " ", " "}, nil}
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

func expressionInBracketsTN(ch []TokenNode) TokenNode {
	return TokenNode{
		Token{ExpressionInBrackets, "expression in brackets", ""}, ch}
}
