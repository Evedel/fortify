package compile

import(
	"io/ioutil"
	"os/exec"
)

import(
	"github.com/Evedel/fortify/src/dictionary"
	"github.com/Evedel/fortify/src/say"
)

func ToFortran(SyntaxTree dictionary.TokenNode, Name [3]string) {
	reducedTree, _ := reduceToFortran(SyntaxTree)
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
