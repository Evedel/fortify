package dictionary

import (
	// "github.com/Evedel/fortify/src/say"
	"strconv"
)

func typeStatic(token Token) (resToken TokenNode) {
	tokval := token.Value

	_, errF := strconv.ParseFloat(tokval, 64)
	if (errF != nil) {
		resToken = TokenNode{Token{String, "String", tokval}, nil}
		return
	} else {
		resToken = TokenNode{Token{Float, "Float", tokval}, nil}
		return
	}
}
