package say

import(
  "github.com/fatih/color"

  // "os"
  "fmt"
  "time"
  "strconv"
)
var mode int
var red     = color.New(color.FgRed).SprintFunc()
var yellow  = color.New(color.FgYellow).SprintFunc()
var cyan    = color.New(color.FgCyan).SprintFunc()


func Init(ll string){
  if ll  == "" {
    mode = 1
    L1("SAY: log level not set. Used default = 1", nil, "\n")
  } else {
    if val, err := strconv.Atoi(ll); err != nil {
      mode = 1
      L1("SAY: log level not int in [1..3]. Used default = 1", nil, "\n")
    } else {
      if (val > -1) && (val < 4) {
        mode = val
      } else {
        mode = 1
        L1("SAY: log level not int in [1..3]. Used default = 1", nil, "\n")
      }
    }
  }
}

func L0(str string, obj interface{}, str2 string){
  fmt.Print(str, obj, str2)
}

func L3(str string, obj interface{}, str2 string){
  // p1 - string, p2 - object, p3 - pretty
  if mode > 0 {
    fmt.Print(red("[ " + time.Now().Format(time.RFC1123) + " ][ L1 ] "),
      red(obj), str2)
  }
}

func L2(str string){
  if mode > 1 {
    fmt.Printf(yellow("[ " + time.Now().Format(time.RFC1123) + " ]" + "[ L2 ] ") + str + "\n")
    fmt.Printf("[ " + time.Now().Format(time.RFC1123) + " ]" + "[ L2 ] " + str + "\n")
  }
}

// func l0(str interface{}){
//   fmt.Print(str)
// }

func L1(str string, obj interface{}, str2 string){
  if mode > 2 {
    fmt.Print(cyan("[ " + time.Now().Format(time.RFC1123) + " ]" + "[ L3 ] "), str,
      obj, str2)
  }
}
