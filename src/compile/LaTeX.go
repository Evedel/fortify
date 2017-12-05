package compile

import (
  "fmt"
  "os/exec"
  "strings"
  "io/ioutil"
)

var Mirrorred = []string{"#"}
var UnMirrorred = []string{"\\say"}

func LaTeX(source string, name [3]string) {
  strsc := source
  strnm := name
  strscnew := ""
  fmt.Print("-> LaTeX compile: ", strnm[0], "\n")
  for _, line := range strings.Split(strsc, "\n") {
    newline := line
    idlc := strings.Index(line, "!")
    if idlc != -1 {
      newline = line[:idlc]
    }
    if idlc != 0 {
      strscnew += newline + "\n"
    }
  }
  strsc = strscnew
  strsc = strings.Replace(strsc, "\n", " \\\\ \n & ", -1)
  for _, ch := range Mirrorred {
    strsc = strings.Replace(strsc, ch, "\\" + string(ch), -1)
  }
  for _, ch := range UnMirrorred {
    strsc = strings.Replace(strsc, ch, "\\text{" + ch[1:] + "}", -1)
  }
  strsc = "& " + strsc + "\n"
  strsc = "\\documentclass[12pt]{article} \n" +
          "\\usepackage{amsmath} \n" +
          "\\begin{document} \n" +
          "\\begin{equation} \n" +
          "\\begin{aligned} \n" +
          strsc +
          "\\end{aligned} \n" +
          "\\end{equation} \n" +
          "\\end{document} \n"

  if err := ioutil.WriteFile("./" + strnm[2] + "/" + strnm[0] + "/" + strnm[0] + ".tex",
    []byte(strsc), 0644); err != nil {
      fmt.Print("-> Error: ", err, "\n")
  } else {
    cmd := exec.Command("pdflatex", "-output-directory", "./" + strnm[2] + "/" + strnm[0] + "/",
      "./" + strnm[2] + "/" + strnm[0] + "/" + strnm[0] + ".tex")
    if output, err := cmd.Output(); err != nil {
      fmt.Print("-> Error: ", err, "\n")
      fmt.Print("-> Error: ", cmd.Args, "\n")
      fmt.Print("-> Error: ", string(output), "\n")
    }
  }
}
