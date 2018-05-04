package dictionary

func TokenNodeProgram() TokenNode {
	return TokenNode{Token{Program, "program", ""}, nil}
}

func TokenNodeExpression() TokenNode {
	return TokenNode{Token{Expression, "expression", ""}, nil}
}

func TokenNodeSpace() TokenNode {
	return TokenNode{Token{Space, "\" \"", " "}, nil}
}

func TokenNodeReturn() TokenNode {
	return TokenNode{Token{CarriageReturn, "\\n", ""}, nil}
}

func TokenNodeVarDec() TokenNode {
  return TokenNode{Token{DeclarationVar, "var", ""}, nil}
}

func TokenNodeVarId(val string) TokenNode {
  return TokenNode{Token{VariableId, "id", val}, nil}
}

func TokenNodeCommentAll() TokenNode {
  return TokenNode{Token{CommentAll, "#", ""}, nil}
}

func TokenNodeString(val string) TokenNode {
  return TokenNode{Token{String, "string", val}, nil}
}

func TokenNodeAssignment() TokenNode {
	return TokenNode{Token{Assignment, "=", ""}, nil}
}

func TokenNodeOperand() TokenNode {
	return TokenNode{Token{Operand, "operand", ""}, nil}
}

func TokenNodeLeftHS() TokenNode {
	return TokenNode{Token{LeftHS, "LeftHS", ""}, nil}
}
func TokenNodeRightHS() TokenNode {
	return TokenNode{Token{RightHS, "RightHS", ""}, nil}
}
func TokenNodeRHS2LHS(t TokenNode) TokenNode {
	return TokenNode{Token{LeftHS, "LeftHS", ""}, t.List}
}

func TokenNodeFromToken(t Token) TokenNode {
	return TokenNode{Token{t.Id, t.IdName, t.Value}, nil}
}
func TokenNodeFromTokenAndList(t Token, tlist []TokenNode) TokenNode {
	return TokenNode{Token{t.Id, t.IdName, t.Value}, tlist}
}
