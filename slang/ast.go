package slang

//A Node is the root of all ast nodes
type Node interface{}

//An Expr represents Expression in slang
type Expr struct{}

//IExpr is a marker interface for Expressions
type IExpr interface {
	Node
	Expr()
}

//Expr Method  implementation to mark Expr as IExpr compliant
func (expr *Expr) Expr() {}

//A NumericConst represents a numeric constant
//Example : 10, 10.23
type NumericConst struct {
	Value float64
}

//Expr Method  implementation to mark NumericConstant as IExpr compliant
func (expr *NumericConst) Expr() {}

//A UnaryExpr node represents a UnaryExpression
//Example : +10, -21.4, -(3 + 2)
type UnaryExpr struct {
	Operator  Token
	RightExpr IExpr
}

//Expr Method  implementation to mark UnaryExpr as IExpr compliant
func (expr *UnaryExpr) Expr() {}

//A BinaryExpr node represents a BinaryExpression
//Example : 10 + 20, 10 * 30 + -20 + 4 / 2
type BinaryExpr struct {
	LeftExpr  IExpr
	Operator  Token
	RightExpr IExpr
}

//Expr Method  implementation to mark BinaryExpr as IExpr compliant
func (expr *BinaryExpr) Expr() {}
