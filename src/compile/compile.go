package compile

import (
  "say"
  "dictionary"

  "os/exec"
  "io/ioutil"
)

func toLatexLang(SyntaxTree dictionary.TokenNode) (Result string) {
  tn := SyntaxTree.This
  if (tn.Id != dictionary.Expression)  &&
      (tn.Id != dictionary.Program)    &&
      (tn.Id != dictionary.CommentTex) &&
      (tn.Id != dictionary.CommentAll) &&
      (tn.Id != dictionary.CommentF90) {
    if (tn.Id == dictionary.CarriageReturn) {
      Result += "\\\\ \n&"
    } else if _, ok := dictionary.SpecialSymbolReverse[tn.Id]; ok {
      if _, ok2 := dictionary.NeedbeMerroredReverse[tn.Id]; ok2 {
        Result += "\\text{\\" + tn.IdName + "}"
      } else {
        Result += "\\text{" + tn.IdName + "}"
      }
    } else if _, ok := dictionary.KeyWordRawReverse[tn.Id]; ok {
      if string(tn.IdName[0]) == "\\" {
        Result += "\\text{\\textbf{" + tn.IdName[1:] + "}}"
      } else {
        Result += "\\text{\\textbf{" + tn.IdName + "}}"
      }
    } else if _, ok := dictionary.DataObjectReverse[tn.Id]; ok {
      Result += "\\text{" + tn.ValueStr + "}"
    }
  }
  if (tn.Id != dictionary.CommentTex) && (tn.Id != dictionary.CommentAll) {
    for ttch := range SyntaxTree.List {
      Result += toLatexLang(SyntaxTree.List[ttch])
    }
  }
  return
}

func toCompiledLang(SyntaxTree dictionary.TokenNode, langtype int) (Result string) {
  tn := SyntaxTree
  tnid := tn.This.Id

  if tnid == dictionary.Program {
    for ttch := range SyntaxTree.List {
      Result += toCompiledLang(SyntaxTree.List[ttch], langtype)
    }
  } else if tnid == dictionary.Expression {
    tnchid := tn.List[0].This.Id
    chlist := tn.List[0].List
    if (tnchid == dictionary.CarriageReturn) ||
        (tnchid == dictionary.Space)         ||
        (tnchid == dictionary.CommentAll)    ||
        (tnchid == dictionary.CommentF90) {
      //------------//
      // do nothing //
      //------------//
    } else if tnchid == dictionary.CommentTex {
      for ttch := range tn.List[0].List {
        Result += toCompiledLang(tn.List[0].List[ttch], langtype)
      }
    } else if tnchid == dictionary.Print {
      if langtype == ltfort {
        Result += "\t" + "print*, "
        for i := 1; i < len(chlist)-1; i++ {
          Result += chlist[i].This.ValueStr
        }
        Result += "\n"
      } else if langtype == ltclang {
        for i := 1; i < len(chlist)-1; i++ {
          Result += "\t" + "printf(" + chlist[i].This.ValueStr + ");\n"
        }
      }
    }
  }
  return
}

func ToFortran(SyntaxTree dictionary.TokenNode, Name [3]string) {
  strnm := Name
  say.L1("Fortran compile: ", strnm[0], "\n")
  srcnew := "program main\n" +
            toCompiledLang(SyntaxTree, ltfort) +
            "end program main"

  say.L0(srcnew, "", "\n")
  program := strnm[2] + "/" + strnm[0] + "/" + strnm[0]
  if err := ioutil.WriteFile(program + ".f90",
    []byte(srcnew), 0644); err != nil {
      say.L3("", err, "\n")
  } else {
    cmd := exec.Command("gfortran", "-o", program+"-f90", program + ".f90")
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
            toCompiledLang(SyntaxTree, ltclang) +
            "}\n"

  say.L0(srcnew, "", "\n")
  program := strnm[2] + "/" + strnm[0] + "/" + strnm[0]
  if err := ioutil.WriteFile(program + ".c",
    []byte(srcnew), 0644); err != nil {
      say.L3("", err, "\n")
  } else {
    cmd := exec.Command("gcc", "-o", program+"-c", program + ".c")
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
          "&" + toLatexLang(SyntaxTree) +
          "\\end{aligned} \n" +
          "\\end{equation} \n" +
          "\\end{document} \n"

  if err := ioutil.WriteFile("./" + strnm[2] + "/" + strnm[0] + "/" + strnm[0] + ".tex",
    []byte(srcnew), 0644); err != nil {
      say.L3("", err, "\n")
  } else {
    cmd := exec.Command("pdflatex", "-output-directory", "./" + strnm[2] + "/" + strnm[0] + "/",
      "./" + strnm[2] + "/" + strnm[0] + "/" + strnm[0] + ".tex")
    if output, err := cmd.Output(); err != nil {
      say.L3("", err, "\n")
      say.L3("", cmd.Args, "\n")
      say.L3("", string(output), "\n")
    }
  }
}
