package compile

import(
  "sort"
)

import(
	"github.com/Evedel/fortify/src/dictionary"
	"github.com/Evedel/fortify/src/say"
)

func toFortran(SyntaxTree dictionary.TokenNode) (Result string, resCode int) {
	resCode = dictionary.UndefinedError

	tn := SyntaxTree
	tnid := tn.This.Id

	tnval := tn.This.Value
	tnidnm := tn.This.IdName
	tnlist := tn.List

	if tnid == dictionary.Program {
		var sortedkeys []string
		for k := range dictionary.Variables {
			sortedkeys = append(sortedkeys, k)
		}
		sort.Strings(sortedkeys)
		for i := range sortedkeys {
			varkey := sortedkeys[i]
			vartype := dictionary.Variables[varkey]
			if vartype == dictionary.Int {
				Result += "\tinteger :: " + varkey + "\n"
			} else if vartype == dictionary.Float {
				Result += "\treal(8) :: " + varkey + "\n"
			} else if vartype == dictionary.String {
				Result += "\tcharacter (len=256) :: " + varkey + "\n"
			}
		}
		for ttch := range SyntaxTree.List {
			op, res := toFortran(SyntaxTree.List[ttch])
			if (res == dictionary.Ok) {
				Result += op
			} else {
				resCode = res
				return
			}
		}
		Result = "program main\n" +
			Result +
			"end program main"
		resCode = dictionary.Ok
	} else if tnid == dictionary.DontCompileTex {
		for ttch := range tnlist {
			op, res := toFortran(tnlist[ttch])
			if (res == dictionary.Ok) {
				Result += op
			} else {
				resCode = res
				return
			}
		}
		resCode = dictionary.Ok
	} else if tnid == dictionary.Print {
		Result += "\t" + "print*, "
		for i := 1; i < len(tnlist)-1; i++ {
			if (i != len(tnlist)-2) {
				Result += tnlist[i].This.Value + ", "
			} else {
				Result += tnlist[i].This.Value
			}
		}
		Result += "\n"
		resCode = dictionary.Ok
	} else if tnid == dictionary.Assignment {
		if (len(tnlist) == 2) {
			op1, r1 := toFortran(tnlist[0])
			op2, r2 := toFortran(tnlist[1])
			if (r1 != dictionary.Ok) {
				resCode = r1
				return
			}
			if (r2 != dictionary.Ok) {
				resCode = r2
				return
			}
			Result += "\t" + op1 + " = " + op2 + "\n"
			resCode = dictionary.Ok
		} else {
			resCode = dictionary.NotEnoughArguments
		}
	} else if (tnid == dictionary.Addition) ||
		(tnid == dictionary.Substraction) ||
		(tnid == dictionary.Multiplication) ||
		(tnid == dictionary.Division) {
			if (len(tnlist) == 2) {
				op1, r1 := toFortran(tnlist[0])
				op2, r2 := toFortran(tnlist[1])

				if (r1 != dictionary.Ok) {
					resCode = r1
					return
				}
				if (r2 != dictionary.Ok) {
					resCode = r2
					return
				}
				Result += op1 + " " + tnidnm + " " + op2
				resCode = dictionary.Ok
			} else {
				// TODO add cases to see wich argument is lost
				resCode = dictionary.NotEnoughArguments
			}
	} else if (tnid == dictionary.VariableId) ||
		(tnid == dictionary.Int) {
		Result = tnval
		resCode = dictionary.Ok
	} else if tnid == dictionary.ExpressionInBrackets {
		op, res := toFortran(tnlist[1])
		if (res == dictionary.Ok) {
			Result += "(" + op + ")"
		} else {
			resCode = res
		}
	} else if ((tnid == dictionary.LeftHS) || (tnid == dictionary.RightHS)) {
		for ttch := range tnlist {
			op, res := toFortran(tnlist[ttch])
			if (res == dictionary.Ok) {
				Result += op
			} else {
				resCode = res
				return
			}
		}
		resCode = dictionary.Ok
		return
	} else if (tnid == dictionary.Float) {
		// say.L0("", tn, "\n")
		Result += tnval
		resCode = dictionary.Ok
	} else {
		say.L3("There is no defined Fortran compiler rules for ["+tnidnm+"]", "", "\n")
	}
	return
}
