package compile

import(
	"github.com/Evedel/fortify/src/dictionary"
	"github.com/Evedel/fortify/src/say"
)

func toLatex(SyntaxTree dictionary.TokenNode) (Result string) {
	tn := SyntaxTree
	tnid := tn.This.Id
	tnchlist := tn.List
	tnchval := tn.This.Value
	tnchidnm := tn.This.IdName
	// say.L0("", tn, "\n")
	if tnid == dictionary.Program {
		for ttchi := 0; ttchi < len(tnchlist); ttchi++ {
			// TODO dirty hack that I don't like, but don't want to work with to the moment
			if (tnchlist[ttchi].This.Id == dictionary.CommentAll) ||
				(tnchlist[ttchi].This.Id == dictionary.DontCompileTex) {
				if len(tnchlist) >= ttchi+1 {
					if tnchlist[ttchi+1].This.Id == dictionary.CarriageReturn {
						ttchi++
					}
				}
			} else {
				Result += toLatex(SyntaxTree.List[ttchi])
			}
		}
		Result = "\\documentclass[8pt]{article} \n" +
			"\\usepackage{amsmath} \n" +
			"\\allowdisplaybreaks\n" +
			"\\begin{document} \n" +
			"\\begin{equation} \n" +
			"\\begin{aligned} \n" +
			"&" + Result + "\\\n" +
			"\\end{aligned} \n" +
			"\\end{equation} \n" +
			"\\end{document} \n"
	} else if tnid == dictionary.Space {
			Result += "\\text{ }"
	} else if tnid == dictionary.CarriageReturn {
		Result += "\\\\\n&"
	} else if tnid == dictionary.DontCompileF90 {
		for ttch := range tnchlist {
			Result += toLatex(tnchlist[ttch])
		}
	} else if tnid == dictionary.Print {
		Result += "\\text{\\textbf{print}}\\{\\text{"
		for i := 1; i < len(tnchlist)-1; i++ {
			Result += tnchlist[i].This.Value
		}
		Result += "}\\}"
	} else if tnid == dictionary.DeclarationVar {
		Result += "\\text{\\textbf{var} }"
		for ttch := range tnchlist {
			Result += toLatex(tnchlist[ttch])
		}
	} else if (tnid == dictionary.VariableId) ||
		(tnid == dictionary.Int) {
		Result += "\\text{" + tnchval + "}"
	} else if (tnid == dictionary.Assignment) ||
		(tnid == dictionary.Addition) ||
		(tnid == dictionary.Substraction) ||
		(tnid == dictionary.Multiplication) ||
		(tnid == dictionary.Division) {
		Result += toLatex(tnchlist[0]) + " " + tnchidnm + " " + toLatex(tnchlist[1])
	} else if ((tnid == dictionary.RightHS) || (tnid == dictionary.LeftHS)) {
		for ttch := range tnchlist {
			Result += toLatex(tnchlist[ttch])
		}
	} else if (tnid == dictionary.Float) {
		Result += "\\text{" + tnchval + "}"
	} else if (tnid == dictionary.ExpressionInBrackets) {
		// Result += "\\text{(}" + toLatex(tnchlist[1]) + "\\text{)}"
		say.L3("There is no defined Latex compiler rules for ["+tnchidnm+"]", "", "\n")
	} else {
		say.L3("There is no defined Latex compiler rules for ["+tnchidnm+"]", "", "\n")
	}
	return
}
