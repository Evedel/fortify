package dictionary

const(
  Ok = iota

  AlreadyDeclared = iota

  LostBracket = iota

  NotALanguageKeyWord = iota
  UnexpectableArgument = iota

  UndefinedError = iota
)

var ErrorCodeDefinitions = map[int]string{
  Ok : "Ok",

  AlreadyDeclared: "AlreadyDeclared",
  
  LostBracket: "LostBracket",

  NotALanguageKeyWord: "NotALanguageKeyWord",
  UnexpectableArgument : "UnexpectableArgument",

  UndefinedError: "UndefinedError" }
