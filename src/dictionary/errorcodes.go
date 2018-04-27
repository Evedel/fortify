package dictionary

const (
	Ok = iota

	AlreadyDeclared = iota

	LostBracket = iota

	NotALanguageKeyWord = iota
	UnexpectedArgument  = iota
	NotANumber          = iota
	NotEnoughArguments  = iota

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
	NotEnoughArguments:  "NotEnoughArguments",

	MissedRoundBracketOpen:  "MissedRoundBracketOpen",
	MissedRoundBracketClose: "MissedRoundBracketClose",

	UndefinedError: "UndefinedError",

	TOTOTODO: "TOTOTODO"}
