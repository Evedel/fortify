Main idea -- write math, not code.
Latex format.
Compiled in pdf and fortran.

Internal code notes
---
rule*
always return index of the last element corresponding to that exact rule, __not__ the first element after  

Comments
---
`!` -- working code for tex, comment in f90 (not in scr.f90 at all)  
`%` -- working code for f90, comment in tex (not in scr.tex at all)  
`#` -- comment everywhere (there is no this string in both src.f90, src.tex)  
TODO  
`#!` -- actual comment in f90, that can be seen in src.f90  
`#%` -- comment in tex, working code for f90  

Variable Declaration
---
Before using, variable needs to be declared  
`var a b c`  
`\var a b c`  
TODO need to be more consistent with latex  
`\var{  a b c  
        d e f }`  

Math
---
All the regular math is supported
`+`, `-`, `*`, `/`, `=`, `(`, `)` -- regular math operation

Sub/Super-scripts  
---  
`_{ij}` -- matrix element [i,j]  
`^{234}` -- 234th power of something  
in variable names this sub/sup/erscripts can be used freely as part of the name  
