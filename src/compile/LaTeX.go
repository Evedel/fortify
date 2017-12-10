package compile

import (
  "say"
  "dictionary"

  "os/exec"
  "io/ioutil"
)

func latexy(SyntaxTree dictionary.TokenNode) (Result string) {
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
      Result += latexy(SyntaxTree.List[ttch])
    }
  }
  return
}

func LaTeX(SyntaxTree dictionary.TokenNode, Name [3]string) {

  strnm := Name
  say.L1("LaTeX compile: ", strnm[0], "\n")
  srcnew := "\\documentclass[8pt]{article} \n" +
          "\\usepackage{amsmath} \n" +
          "\\allowdisplaybreaks\n" +
          "\\begin{document} \n" +
          "\\begin{equation} \n" +
          "\\begin{aligned} \n" +
          "&" + latexy(SyntaxTree) +
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
