Main idea -- write math, not code.
Latex format.
Compiled in pdf and fortran.

`var` set of variables (definition)
array assignment `= []`
array read through `A_{i,j,k}`
`\Function` == internal functions (Latex or ff)
`! Expression` == compile in fortran, but not in latex
`# Expression` == comment


Program :=
  1)  `EOF`
  2)  Statement `\n`
  3)  Statement `\n`
      Program

Statement :=
  1) ``
  2) `\n`
  3) ` `{Statement}
  4) OptionAssignment
  5) VariableAssignment
  6) Expression

Expression :=
  1) Condition
  2) MathExpression
  3) Value

VariableAssignment :=               \\ Only previously defined keys
  1) VariableKey `=` Value

RoundParentheses :=
  1) `(` Expression `)`

OptionAssignment :=
  1) OptionKey `=` Value

VariableDefinition :=               \\ Only first definition
  2) `int` VariableKey, {`,` VariableKey}
  3) `float` VariableKey, {`,` VariableKey}
  4) `var` VariableKey, {`,` VariableKey}

VariableKey :=
  1) Character{Character || Digit}

Array :=
  1) `{` Value {`,` Value } `}`
  2) `{` Array {`,` Array } `}`
  3) `{` Number `,` Number `,` Number `}`

Value :=
  1) Number
  2) Array

Number :=
  1) Integer
  2) Float

Integer :=
  1) Digit{Digit}

Float :=
  1) `.`Digit{Digit}
  3) Digit{Digit}`.`Digit{Digit}

Character :=
  1) `a`, `b`, ... , `z`, `A`, `B`, ... `Z`, `_`

Digit :=
  1) `0`, `1`, ... , `9`
