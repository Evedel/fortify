package compile

import (
  "say"
  "dictionary"

  "os/exec"
  "io/ioutil"
)

func f90fy(SyntaxTree dictionary.TokenNode) (Result string) {
  tn := SyntaxTree
  tnid := tn.This.Id

  if tnid == dictionary.Program {
    for ttch := range SyntaxTree.List {
      Result += f90fy(SyntaxTree.List[ttch])
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
        Result += f90fy(tn.List[0].List[ttch])
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

func Fortran(SyntaxTree dictionary.TokenNode, Name [3]string) {
  strnm := Name
  say.L1("Fortran compile: ", strnm[0], "\n")
  srcnew := "program main\n" +
            f90fy(SyntaxTree) +
            "end program main"

  say.L0(srcnew, "", "\n")
  program := strnm[2] + "/" + strnm[0] + "/" + strnm[0]
  if err := ioutil.WriteFile(program + ".f90",
    []byte(srcnew), 0644); err != nil {
      say.L3("", err, "\n")
  } else {
    cmd := exec.Command("gfortran", "-o", program, program + ".f90")
    if output, err := cmd.Output(); err != nil {
      say.L3("", err, "\n")
      say.L3("", cmd.Args, "\n")
      say.L3("", string(output), "\n")
    }
  }
}
