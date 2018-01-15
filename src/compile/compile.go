package compile

import(
	"io/ioutil"
	"os/exec"
	"sort"
)

// import(
// 	"dictionary"
// 	"say"
// )
import(
	"github.com/Evedel/fortify/src/dictionary"
	"github.com/Evedel/fortify/src/say"
)

func toLatex(SyntaxTree dictionary.TokenNode) (Result string) {
	tn := SyntaxTree
	tnid := tn.This.Id
	if tnid == dictionary.Program {
		for ttch := range SyntaxTree.List {
			Result += toLatex(SyntaxTree.List[ttch])
		}
	} else if tnid == dictionary.Expression {
		for i := range tn.List {
			tnchid := tn.List[i].This.Id
			tnchval := tn.List[i].This.Value
			tnchidnm := tn.List[i].This.IdName
			tnchlist := tn.List[i].List

			if (tnchid == dictionary.CommentAll) ||
				(tnchid == dictionary.CommentTex) {
				//------------//
				// do nothing //
				//------------//
			} else if tnchid == dictionary.Space {
				Result += "\\text{ }"
			} else if tnchid == dictionary.CarriageReturn {
				Result += "\\\\\n&"
			} else if tnchid == dictionary.CommentF90 {
				for ttch := range tnchlist {
					Result += toLatex(tnchlist[ttch])
				}
			} else if tnchid == dictionary.Print {
				Result += "\\text{\\textbf{print}}\\{\\text{"
				for i := 1; i < len(tnchlist)-1; i++ {
					Result += tnchlist[i].This.Value
				}
				Result += "}\\}"
			} else if tnchid == dictionary.DeclarationVar {
				Result += "\\text{\\textbf{var} }"
				for ttch := range tnchlist {
					Result += toLatex(tnchlist[ttch])
				}
			} else if (tnchid == dictionary.VariableId) ||
				(tnchid == dictionary.Int) {
				Result += "\\text{" + tnchval + "}"
			} else if (tnchid == dictionary.Assignment) ||
				(tnchid == dictionary.Addition) ||
				(tnchid == dictionary.Substraction) ||
				(tnchid == dictionary.Multiplication) ||
				(tnchid == dictionary.Division) {
				Result += toLatex(tnchlist[0]) + " " + tnchidnm + " " + toLatex(tnchlist[1])
			} else if tnchid == dictionary.ExpressionInBrackets {
				Result += "\\text{(}" + toLatex(tnchlist[1]) + "\\text{)}"
			} else {
				say.L3("There is no defined Latex compiler rules for ["+tnchidnm+"]", "", "\n")
			}
		}
	}
	return
}

func toFortran(SyntaxTree dictionary.TokenNode) (Result string) {
	tn := SyntaxTree
	tnid := tn.This.Id
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
			}
			if vartype == dictionary.Float {
				Result += "\treal(8) :: " + varkey + "\n"
			}
			if vartype == dictionary.String {
				Result += "\tcharacter (len=256) :: " + varkey + "\n"
			}
		}
		for ttch := range SyntaxTree.List {
			Result += toFortran(SyntaxTree.List[ttch])
		}
	} else if tnid == dictionary.Expression {
		for i := range tn.List {
			tnchid := tn.List[i].This.Id
			tnchval := tn.List[i].This.Value
			tnchidnm := tn.List[i].This.IdName
			tnchlist := tn.List[i].List

			if (tnchid == dictionary.CarriageReturn) ||
				(tnchid == dictionary.Space) ||
				(tnchid == dictionary.CommentAll) ||
				(tnchid == dictionary.CommentF90) ||
				(tnchid == dictionary.DeclarationVar) {
				//------------//
				// do nothing //
				//------------//
			} else if tnchid == dictionary.CommentTex {
				for ttch := range tnchlist {
					Result += toFortran(tnchlist[ttch])
				}
			} else if tnchid == dictionary.Print {
				Result += "\t" + "print*, "
				for i := 1; i < len(tnchlist)-1; i++ {
					if tnchlist[i].This.Id == dictionary.Space {
						Result += ", "
					} else {
						Result += tnchlist[i].This.Value
					}
				}
				Result += "\n"
			} else if tnchid == dictionary.Assignment {
				Result += "\t" + toFortran(tnchlist[0]) + " = " + toFortran(tnchlist[1]) + "\n"
			} else if (tnchid == dictionary.Addition) ||
				(tnchid == dictionary.Substraction) ||
				(tnchid == dictionary.Multiplication) ||
				(tnchid == dictionary.Division) {
				Result += toFortran(tnchlist[0]) + " " + tnchidnm + " " + toFortran(tnchlist[1])
			} else if (tnchid == dictionary.VariableId) ||
				(tnchid == dictionary.Int) {
				Result = tnchval
			} else if tnchid == dictionary.ExpressionInBrackets {
				Result += "(" + toFortran(tnchlist[1]) + ")"
			} else {
				say.L3("There is no defined Fortran compiler rules for ["+tnchidnm+"]", "", "\n")
			}
		}
	}
	return
}

func toClang(SyntaxTree dictionary.TokenNode) (Result string) {
	tn := SyntaxTree
	tnid := tn.This.Id

	if tnid == dictionary.Program {
		for ttch := range SyntaxTree.List {
			Result += toClang(SyntaxTree.List[ttch])
		}
	} else if tnid == dictionary.Expression {
		tnchid := tn.List[0].This.Id
		chlist := tn.List[0].List
		if (tnchid == dictionary.CarriageReturn) ||
			(tnchid == dictionary.Space) ||
			(tnchid == dictionary.CommentAll) ||
			(tnchid == dictionary.CommentF90) {
			//------------//
			// do nothing //
			//------------//
		} else if tnchid == dictionary.CommentTex {
			for ttch := range tn.List[0].List {
				Result += toClang(tn.List[0].List[ttch])
			}
		} else if tnchid == dictionary.Print {
			for i := 1; i < len(chlist)-1; i++ {
				Result += "\t" + "printf(" + chlist[i].This.Value + ");\n"
			}
		}
	}
	return
}

func ToFortran(SyntaxTree dictionary.TokenNode, Name [3]string) {
	// dictionary.PrintSyntaxTree(SyntaxTree, "")
	strnm := Name
	say.L1("Fortran compile: ", strnm[0], "\n")
	srcnew := "program main\n" +
		toFortran(SyntaxTree) +
		"end program main"

	say.L0(srcnew, "", "\n")
	program := strnm[2] + "/" + strnm[0] + "/" + strnm[0]
	if err := ioutil.WriteFile(program+".f90",
		[]byte(srcnew), 0644); err != nil {
		say.L3("", err, "\n")
	} else {
		cmd := exec.Command("gfortran", "-o", program, program+".f90")
		if output, err := cmd.Output(); err != nil {
			say.L3("", err, "\n")
			say.L3("", cmd.Args, "\n")
			say.L3("", string(output), "\n")
		}
	}
}

func ToClang(SyntaxTree dictionary.TokenNode, Name [3]string) {
	strnm := Name
	say.L1("Clang compile: ", strnm[0], "\n")
	srcnew := "#include<stdio.h>\n" +
		"main()\n" +
		"{\n" +
		toClang(SyntaxTree) +
		"}\n"

	say.L0(srcnew, "", "\n")
	program := strnm[2] + "/" + strnm[0] + "/" + strnm[0]
	if err := ioutil.WriteFile(program+".c",
		[]byte(srcnew), 0644); err != nil {
		say.L3("", err, "\n")
	} else {
		cmd := exec.Command("gcc", "-o", program+"-c", program+".c")
		if output, err := cmd.Output(); err != nil {
			say.L3("", err, "\n")
			say.L3("", cmd.Args, "\n")
			say.L3("", string(output), "\n")
		}
	}
}

func ToLaTeX(SyntaxTree dictionary.TokenNode, Name [3]string) {
	strnm := Name
	say.L1("LaTeX compile: ", strnm[0], "\n")
	srcnew := "\\documentclass[8pt]{article} \n" +
		"\\usepackage{amsmath} \n" +
		"\\allowdisplaybreaks\n" +
		"\\begin{document} \n" +
		"\\begin{equation} \n" +
		"\\begin{aligned} \n" +
		"&" + toLatex(SyntaxTree) +
		"\\end{aligned} \n" +
		"\\end{equation} \n" +
		"\\end{document} \n"

	if err := ioutil.WriteFile("./"+strnm[2]+"/"+strnm[0]+"/"+strnm[0]+".tex",
		[]byte(srcnew), 0644); err != nil {
		say.L3("", err, "\n")
	} else {
		cmd := exec.Command("pdflatex", "-output-directory", "./"+strnm[2]+"/"+strnm[0]+"/",
			"./"+strnm[2]+"/"+strnm[0]+"/"+strnm[0]+".tex")
		if output, err := cmd.Output(); err != nil {
			say.L3("", err, "\n")
			say.L3("", cmd.Args, "\n")
			say.L3("", string(output), "\n")
		}
	}
}
