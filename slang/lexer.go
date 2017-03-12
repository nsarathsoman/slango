package slang

import (
	"fmt"
	"strconv"
)

//Lexer represents the lexical analyzer
type Lexer struct {
	Data      []byte
	DataLen   int
	Index     int
	CurToken  Token
	PrevToken Token
	Num       float64
}

//Reads the data stream and fetches the token
func (lexer *Lexer) eat() {

	//Set to unknown if end of stream is found
	if lexer.IsEndOfStream() {
		lexer.PrevToken = lexer.CurToken
		lexer.CurToken = UNKNOWN
	}

dataStream:
	for lexer.IsNotEndOfStream() {
		switch lexer.Data[lexer.Index] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			lexer.readNum()
			lexer.PrevToken = lexer.CurToken
			lexer.CurToken = NUM
			break dataStream
		case '+':
			lexer.Index++
			lexer.PrevToken = lexer.CurToken
			lexer.CurToken = ADD
			break dataStream
		case '-':
			lexer.Index++
			lexer.PrevToken = lexer.CurToken
			lexer.CurToken = SUB
			break dataStream
		case '/':
			lexer.Index++
			lexer.PrevToken = lexer.CurToken
			lexer.CurToken = DIV
			break dataStream
		case '*':
			lexer.Index++
			lexer.PrevToken = lexer.CurToken
			lexer.CurToken = MUL
			break dataStream
		case ' ', '\n':
			lexer.Index++
		case '(':
			lexer.Index++
			lexer.PrevToken = lexer.CurToken
			lexer.CurToken = OPAR
			break dataStream
		case ')':
			lexer.Index++
			lexer.PrevToken = lexer.CurToken
			lexer.CurToken = CPAR
			break dataStream
		default:
			lexer.PrevToken = lexer.CurToken
			lexer.CurToken = UNKNOWN
			break dataStream
			// panic(fmt.Sprintf("Unknown character at %v", lexer.Index+1))
		}
	}

}

//reads number from the source data stream
func (lexer *Lexer) readNum() {
	var foundDot bool
	var data = make([]byte, 0)

	for lexer.IsNotEndOfStream() && lexer.isNumeric() {
		if '.' == lexer.Data[lexer.Index] {
			if !foundDot {
				foundDot = true
			} else {
				panic("While parsing numeric value Found '.' more than two times")
			}
		}

		data = append(data, lexer.Data[lexer.Index])
		lexer.Index++
	}
	num, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		panic("Unable to convert string to number while lexical analysis")
	}
	lexer.Num = num
}

//check for end of stream
func (lexer *Lexer) IsNotEndOfStream() bool {
	if lexer.DataLen <= lexer.Index {
		return false
	}
	return true

}

//check for not end of stream
func (lexer *Lexer) IsEndOfStream() bool {
	return !lexer.IsNotEndOfStream()

}

//isNumeric checks whether a char is numeric or not
func (lexer *Lexer) isNumeric() bool {
	if lexer.IsEndOfStream() {
		panic("Cannot check for numeric char; End of stream reached")
	}

	switch lexer.Data[lexer.Index] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
		return true
	default:
		return false
	}
}

func (lexer *Lexer) expect(expectedToken Token) {
	actualToken := lexer.CurToken
	if byte(actualToken) != byte(expectedToken) {
		panic(fmt.Sprintf("Expecting a %v", expectedToken))
	}
}
