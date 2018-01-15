package dictionary

const (
	Ok = iota

	AlreadyDeclared = iota

	LostBracket = iota

	NotALanguageKeyWord = iota
	UnexpectedArgument  = iota
	NotANumber          = iota

	MissedRoundBracketClose = iota
	MissedRoundBracketOpen  = iota

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

	MissedRoundBracketOpen:  "MissedRoundBracketOpen",
	MissedRoundBracketClose: "MissedRoundBracketClose",

	UndefinedError: "UndefinedError",

	TOTOTODO: "TOTOTODO"}
