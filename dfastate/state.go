package dfastate

type State string

const (
	Init       = "Init"
	Id         = "Id"
	GT         = ">"
	GE         = ">="
	IntLiteral = "IntLiteral"
	Plus       = "+"
	Assign     = "Assign"
	EQ         = "EQ"
)
