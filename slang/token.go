package slang

type Token byte

const (
	//Arithmetic Operators
	ADD Token = Token(1) // + operator
	SUB       = Token(2) // - operator
	DIV       = Token(3) // '/' operator
	MUL       = Token(4) // *operator

	OPAR = Token(5) //Open paranthesis '('
	CPAR = Token(6) //Closed paranthesis ')'

	NUM = Token(7)

	UNKNOWN = Token(8)
)
