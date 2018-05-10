package compile

import(
	"io/ioutil"
	"os/exec"
	"sort"
)

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
		for ttch := range SyntaxTree.List {
			Result += toLatex(SyntaxTree.List[ttch])
		}
		Result = "\\documentclass[8pt]{article} \n" +
			"\\usepackage{amsmath} \n" +
			"\\allowdisplaybreaks\n" +
			"\\begin{document} \n" +
			"\\begin{equation} \n" +
			"\\begin{aligned} \n" +
			"&" + Result +
			"\\end{aligned} \n" +
			"\\end{equation} \n" +
			"\\end{document} \n"
	} else if (tnid == dictionary.CommentAll) ||
		(tnid == dictionary.DontCompileTex) {
		//------------//
		// do nothing //
		//------------//
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
				Result += "\t" + op1 + " " + tnidnm + " " + op2 + "\n"
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
		op, res := toFortran(tnlist[0])
		if (res == dictionary.Ok) {
			Result += op
			resCode = dictionary.Ok
		} else {
			resCode = res
		}
	} else if (tnid == dictionary.Float) {
		Result += tnval
		resCode = dictionary.Ok
	} else {
		say.L3("There is no defined Fortran compiler rules for ["+tnidnm+"]", "", "\n")
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
	} else {
		tnchid := tn.List[0].This.Id
		chlist := tn.List[0].List
		if (tnchid == dictionary.CarriageReturn) ||
			(tnchid == dictionary.Space) ||
			(tnchid == dictionary.CommentAll) ||
			(tnchid == dictionary.DontCompileF90) {
			//------------//
			// do nothing //
			//------------//
		} else if tnchid == dictionary.DontCompileTex {
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
	reducedTree, _ := reduceToFortran(SyntaxTree)
	// dictionary.PrintSyntaxTree(reducedTree, "")
	strnm := Name
	say.L1("Fortran compile: ", strnm[0], "\n")
	src, res := toFortran(reducedTree)
	srcend := ""
	if (res == dictionary.Ok) {
		srcend = src
	} else {
		say.L1("", src, "\n")
		say.L1("", dictionary.ErrorCodeDefinitions[res], "\n")
		return
	}

	say.L0(srcend, "", "\n")
	program := strnm[2] + "/" + strnm[0] + "/" + strnm[0]
	if err := ioutil.WriteFile(program+".f90",
		[]byte(srcend), 0644); err != nil {
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
	if err := ioutil.WriteFile("./"+strnm[2]+"/"+strnm[0]+"/"+strnm[0]+".tex",
		[]byte(toLatex(SyntaxTree)), 0644); err != nil {
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
