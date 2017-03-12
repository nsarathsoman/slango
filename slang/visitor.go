package slang

//Visitor is
type Visitor interface {
	Visit(n Node) *NumericConst
}

//Interpreter implementation
type Interpreter struct{}

//Visit method of Visiotor is attached up on interface
func (i *Interpreter) Visit(n Node) *NumericConst {
	switch ast := n.(type) {
	case *BinaryExpr:
		leftValue := i.Visit(ast.LeftExpr)
		rightValue := i.Visit(ast.RightExpr)

		switch ast.Operator {
		case ADD:
			return &NumericConst{Value: leftValue.Value + rightValue.Value}
		case SUB:
			return &NumericConst{Value: leftValue.Value - rightValue.Value}
		case DIV:
			return &NumericConst{Value: leftValue.Value / rightValue.Value}
		case MUL:
			return &NumericConst{Value: leftValue.Value * rightValue.Value}
		default:
			panic("Unknown Operator")
		}

	case *UnaryExpr:
		rightValue := i.Visit(ast.RightExpr)
		switch ast.Operator {
		case ADD:
			return &NumericConst{Value: 0 + rightValue.Value}
		case SUB:
			return &NumericConst{Value: 0 - rightValue.Value}
		default:
			panic("Unknown Operator")
		}

	case *NumericConst:
		return ast

	default:
		panic("Unknown Node")
	}
}
