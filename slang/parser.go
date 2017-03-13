package slang

import "io/ioutil"

//Struct Parser represents the parser component
type Parser struct {
	ModulePath string
	Lexer      *Lexer
}

//parse inititates a module parsing
func parse(modulePath string) IExpr {
	data, err := ioutil.ReadFile(modulePath)
	if err != nil {
		panic(err)
	}
	return ParseFromStream(data)
}

func ParseFromStream(data []byte) IExpr {
	parser := &Parser{}
	parser.Lexer = &Lexer{Data: data, DataLen: len(data), Index: 0}
	expr := parser.parseExpr()
	return expr
}

//{Expr} := {Term} || {Term} {'+' | '-'} {Expr}
func (parser *Parser) parseExpr() IExpr {
	lexer := parser.Lexer
	expr := parser.parseTerm()
	// lexer.eat()
	switch token := lexer.CurToken; token {
	case ADD, SUB:
		rightExp := parser.parseExpr()
		expr = &BinaryExpr{LeftExpr: expr, Operator: token, RightExpr: rightExp}

	}
	return expr
}

//{Term} := {Factor} || {Factor} {'*' | '/'} {Term}
func (parser *Parser) parseTerm() IExpr {
	lexer := parser.Lexer
	//operator predecense for division
	expr := parser.parseFactor()
	lexer.eat()
	token := lexer.CurToken
	for token == DIV || token == MUL {
		if lexer.CurToken == DIV {
			rightExp := parser.parseFactor()
			expr = &BinaryExpr{LeftExpr: expr, Operator: token, RightExpr: rightExp}
			lexer.eat()
			token = lexer.CurToken
		} else {
			rightExp := parser.parseTerm()
			expr = &BinaryExpr{LeftExpr: expr, Operator: token, RightExpr: rightExp}
			token = lexer.CurToken
		}
	}
	return expr
}

//{Factor} := {Num} || '(' {Expr} ')' {'+' | '-'} {Factor}
func (parser *Parser) parseFactor() IExpr {
	lexer := parser.Lexer
	lexer.eat()
	switch token := lexer.CurToken; token {
	case NUM:
		return &NumericConst{Value: lexer.Num}
	case ADD:
		expr := parser.parseFactor()
		return expr
	case SUB:
		expr := parser.parseFactor()
		return &UnaryExpr{Operator: SUB, RightExpr: expr}
	case OPAR:
		expr := parser.parseExpr()
		lexer.expect(CPAR)
		return expr
	default:
		panic("Unknown production")
		// return parser.parseExpr()

	}
}
