package dictionary

const (
	Ok = iota

	AlreadyDeclared = iota

	LostBracket = iota

	NotALanguageKeyWord = iota
	UnexpectedArgument  = iota
	NotANumber          = iota

	UndefinedError = iota

	TOTOTODO = iota
)

var ErrorCodeDefinitions = map[int]string{
	Ok: "Ok",

	AlreadyDeclared: "AlreadyDeclared",

	LostBracket: "LostBracket",

	NotALanguageKeyWord: "NotALanguageKeyWord",
	UnexpectedArgument:  "UnexpectedArgument",
	NotANumber:          "NotANumber",

	UndefinedError: "UndefinedError",

	TOTOTODO: "TOTOTODO"}
