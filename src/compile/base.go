package compile

import (
  // "os/exec"
  "strings"
  "io/ioutil"
)

// str := "program main \n"
// str += "implicit none \n"
// str += "print*, \"hello\"\n"
// str += "end program\n"

// _ = ioutil.WriteFile("./tmp.f90", []byte(str), 0644)
// cmd := exec.Command("gfortran", "-v", "-o", "tmp", "tmp.f90")
// _ = cmd.Run()
// cmd := exec.Command("gfortran", "-v", "-o", "tmp", "tmp.f90")
// _ = cmd.Run()

func LaTeX(source string, name string) {
  strsc := source
  strnm := name
  strsc = strings.Replace(strsc, "\n", " \\\\ \n", -1)
  strsc = "\\documentclass[12pt]{article} \n" +
          "\\usepackage{amsmath} \n" +
          "\\begin{document} \n" +
          "\\begin{equation} \n" +
          "\\begin{aligned} \n" +
          strsc +
          "\\end{aligned} \n" +
          "\\end{equation} \n" +
          "\\end{document} \n"


  i := strings.LastIndex(strnm, ".")
  if i != -1 {
    strnm = strnm[:i]
  }
  _ = ioutil.WriteFile("./" + strnm + ".tex", []byte(strsc), 0644)
}
