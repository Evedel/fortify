package compile

import(
	"github.com/Evedel/fortify/src/dictionary"
)

func toClang(SyntaxTree dictionary.TokenNode) (Result string) {
	tn := SyntaxTree
	tnid := tn.This.Id

	if tnid == dictionary.Program {
		for ttch := range SyntaxTree.List {
			Result += toClang(SyntaxTree.List[ttch])
		}
	} else {
		tnchid := tn.List[0].This.Id
		chlist := tn.List[0].List
		if (tnchid == dictionary.CarriageReturn) ||
			(tnchid == dictionary.Space) ||
			(tnchid == dictionary.CommentAll) ||
			(tnchid == dictionary.DontCompileF90) {
			//------------//
			// do nothing //
			//------------//
		} else if tnchid == dictionary.DontCompileTex {
			for ttch := range tn.List[0].List {
				Result += toClang(tn.List[0].List[ttch])
			}
		} else if tnchid == dictionary.Print {
			for i := 1; i < len(chlist)-1; i++ {
				Result += "\t" + "printf(" + chlist[i].This.Value + ");\n"
			}
		}
	}
	return
}
