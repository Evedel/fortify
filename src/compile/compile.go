package compile

import (
  "say"
  "dictionary"

  "os/exec"
  "io/ioutil"
)

func toLatex(SyntaxTree dictionary.TokenNode) (Result string) {
  tn := SyntaxTree
  tnid := tn.This.Id
  if tnid == dictionary.Program {
    for ttch := range SyntaxTree.List {
      Result += toLatex(SyntaxTree.List[ttch])
    }
  } else if tnid == dictionary.Expression {
    tnchid := tn.List[0].This.Id
    tnchlist := tn.List[0].List
    tnchidnm := tn.List[0].This.IdName

    if (tnchid == dictionary.CommentAll)  ||
        (tnchid == dictionary.CommentTex) {
      //------------//
      // do nothing //
      //------------//
    } else if tnchid == dictionary.CarriageReturn {
      Result += "\\\\\n&"
    } else if tnchid == dictionary.CommentF90 {
      for ttch := range tnchlist {
        Result += toLatex(tnchlist[ttch])
      }
    } else if _, ok := dictionary.SpecialSymbolReverse[tnchid]; ok {
      if _, ok2 := dictionary.NeedbeMerroredReverse[tnchid]; ok2 {
        Result += "\\text{\\" + tnchidnm + "}"
      } else {
        Result += "\\text{" + tnchidnm + "}"
      }
    } else if tnchid == dictionary.Print {
      Result += "\\text{\\textbf{print}}\\{"
      for i := 1; i < len(tnchlist)-1; i++ {
        Result += "\\text{" + tnchlist[i].This.ValueStr + "}"
      }
      Result += "\\}"
    }
  }
  return
}

func toFortran(SyntaxTree dictionary.TokenNode) (Result string) {
  tn := SyntaxTree
  tnid := tn.This.Id

  if tnid == dictionary.Program {
    for ttch := range SyntaxTree.List {
      Result += toFortran(SyntaxTree.List[ttch])
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
        Result += toFortran(tn.List[0].List[ttch])
      }
    } else if tnchid == dictionary.Print {
      Result += "\t" + "print*, "
      for i := 1; i < len(chlist)-1; i++ {
        Result += chlist[i].This.ValueStr
      }
      Result += "\n"
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
        (tnchid == dictionary.Space)         ||
        (tnchid == dictionary.CommentAll)    ||
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
        Result += "\t" + "printf(" + chlist[i].This.ValueStr + ");\n"
      }
    }
  }
  return
}

func ToFortran(SyntaxTree dictionary.TokenNode, Name [3]string) {
  strnm := Name
  say.L1("Fortran compile: ", strnm[0], "\n")
  srcnew := "program main\n" +
            toFortran(SyntaxTree) +
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
            toClang(SyntaxTree) +
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
          "&" + toLatex(SyntaxTree) +
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
