package compile

import(
	"github.com/Evedel/fortify/src/dictionary"
	"github.com/Evedel/fortify/src/say"
)

func toLatex(SyntaxTree dictionary.TokenNode) (Result string) {
	tn := SyntaxTree
	tnid := tn.This.Id
	tnlist := tn.List
	tnval := tn.This.Value
	tnidnm := tn.This.IdName
	// say.L0("", tn, "\n")
	if tnid == dictionary.Program {
		for tti := 0; tti < len(tnlist); tti++ {
			// TODO dirty hack that I don't like, but don't want to work with to the moment
			if (tnlist[tti].This.Id == dictionary.CommentAll) ||
				(tnlist[tti].This.Id == dictionary.DontCompileTex) {
				if len(tnlist) >= tti+1 {
					if tnlist[tti+1].This.Id == dictionary.CarriageReturn {
						tti++
					}
				}
			} else {
				Result += toLatex(SyntaxTree.List[tti])
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
		for tt := range tnlist {
			Result += toLatex(tnlist[tt])
		}
	} else if tnid == dictionary.Print {
		Result += "\\text{\\textbf{print}}\\{\\text{"
		for i := 1; i < len(tnlist)-1; i++ {
			Result += tnlist[i].This.Value
		}
		Result += "}\\}"
	} else if tnid == dictionary.DeclarationVar {
		Result += "\\text{\\textbf{var} }"
		for tt := range tnlist {
			Result += toLatex(tnlist[tt])
		}
	} else if (tnid == dictionary.VariableId) ||
		(tnid == dictionary.Int) {
		Result += "\\text{" + tnval + "}"
	} else if (tnid == dictionary.Assignment) ||
		(tnid == dictionary.Addition) ||
		(tnid == dictionary.Substraction) ||
		(tnid == dictionary.Multiplication) ||
		(tnid == dictionary.Division) {
		Result += toLatex(tnlist[0]) + " " + tnidnm + " " + toLatex(tnlist[1])
	} else if ((tnid == dictionary.RightHS) || (tnid == dictionary.LeftHS)) {
		for tt := range tnlist {
			Result += toLatex(tnlist[tt])
		}
	} else if (tnid == dictionary.Float) {
		Result += "\\text{" + tnval + "}"
	} else if (tnid == dictionary.RoundBrackets) {
		if len(tnlist) != 0 {
			Result += "\\text{(}" + toLatex(tnlist[0]) + "\\text{)}"
		} else {
			Result += "\\text{(}\\text{)}"
		}
	} else {
		say.L3("There is no defined Latex compiler rules for ["+tnidnm+"]", "", "\n")
	}
	return
}
