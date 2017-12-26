package dictionary

import (
	// "github.com/Evedel/fortify/src/say"
	"strconv"
)

func typeNumber(ttail []Token) (resCode int, stopInd int, resToken TokenNode, errmsg string) {
	resCode = UndefinedError
	stopInd = 0

	tokval := ttail[0].Value

	_, errF := strconv.ParseFloat(tokval, 64)
	_, errI := strconv.ParseInt(tokval, 10, 64)
	if (errF != nil) && (errI != nil) {
		resCode = NotANumber
		return
	} else {
		if errF == nil {
			resToken = TokenNode{Token{Float, "Float", tokval}, nil}
		}
		if errI == nil {
			resToken = TokenNode{Token{Int, "Int", tokval}, nil}
		}
		resCode = Ok
	}
	return
}
