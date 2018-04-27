package dictionary

// all the possible words
const (
	// Internal tokens
	None        = iota
	Program    						= iota
	Expression 						= iota
	ExpressionInBrackets 	= iota
	VariableId = iota

	OperatonUnary   = iota
	OperatonBinary  = iota
	OperatorTernary = iota
	Operand  				= iota
	// Symbols
	CarriageReturn = iota
	DoubleQuote    = iota
	Comma          = iota
	Space          = iota
	CommentTex = iota // Not shown in tex, but compiled in f90
	CommentF90 = iota // Not shown in f90, but compiled in tex
	CommentAll = iota // Not compiled at all
	CurlyBracketOpen  = iota
	CurlyBracketClose = iota
	RoundBracketOpen  = iota
	RoundBracketClose = iota
	// Binary operators
	Assignment     = iota
	Addition       = iota
	Substraction   = iota
	Multiplication = iota
	Division       = iota
	// Data types
	Int    = iota
	Float  = iota
	Word   = iota
	String = iota
	// Key words
	Print          = iota
	DeclarationVar = iota
)

var SpecialSymbolReverse = map[int]string{}
var SpecialSymbol = map[string]int{
	"\n": CarriageReturn,
	"%":  CommentTex,
	"!":  CommentF90,
	"#":  CommentAll,
	"\"": DoubleQuote,
	"{":  CurlyBracketOpen,
	"}":  CurlyBracketClose,
	",":  Comma,
	" ":  Space,
	"=":  Assignment,
	"+":  Addition,
	"-":  Substraction,
	"*":  Multiplication,
	"/":  Division,
	"(":  RoundBracketOpen,
	")":  RoundBracketClose}

var NeedbeMerroredReverse = map[int]string{
	CurlyBracketOpen:  "{",
	CurlyBracketClose: "}",
	CommentAll:        "#"}

var KeyWordRawReverse = map[int]string{}
var KeyWordRaw = map[string]int{
	"print": Print,
	"var ":  DeclarationVar}

var KeyWordBackslashReverse = map[int]string{}
var KeyWordBackslash = map[string]int{
	"\\print": Print,
	"\\var":   DeclarationVar}

var DataObjectReverse = map[int]string{}
var DataObject = map[string]int{
	"word":   Word,
	"string": String}

var Variables = map[string]int{}

type Token struct {
	Id     int
	IdName string
	Value  string
}

type TokenNode struct {
	This Token
	List []TokenNode
}
