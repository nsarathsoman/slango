package slang

import (
	"fmt"
	"testing"
)

func TestInterpretter(t *testing.T) {
	//test case 1
	t.Run("case1", func(t *testing.T) {
		expr := &BinaryExpr{
			LeftExpr: &NumericConst{Value: 20.0},
			Operator: ADD,
			RightExpr: &UnaryExpr{
				Operator:  SUB,
				RightExpr: &NumericConst{Value: 10.99}}}
		testInterpretter(expr, 9.01, t)
	})

	//test case 2
	t.Run("case2", func(t *testing.T) {
		expr := &BinaryExpr{
			LeftExpr: &NumericConst{Value: 20.0},
			Operator: MUL,
			RightExpr: &UnaryExpr{
				Operator:  ADD,
				RightExpr: &NumericConst{Value: 10.00}}}
		testInterpretter(expr, 200.00, t)
	})

	//test case 3
	t.Run("case3", func(t *testing.T) {
		expr := &BinaryExpr{
			LeftExpr: &NumericConst{Value: 20.0},
			Operator: DIV,
			RightExpr: &UnaryExpr{
				Operator:  SUB,
				RightExpr: &NumericConst{Value: 10.00}}}
		testInterpretter(expr, -2.0, t)
	})

	//test case 4
	t.Run("case4", func(t *testing.T) {
		expr := &BinaryExpr{
			LeftExpr: &NumericConst{Value: 20.0},
			Operator: DIV,
			RightExpr: &UnaryExpr{
				Operator:  MUL,
				RightExpr: &NumericConst{Value: 10.00}}}
		defer func() {
			if rec := recover(); rec == nil {
				t.Error("Interpretter not throwing error " +
					"when unsupported operator is applied with unary expression")
			}
		}()
		testInterpretter(expr, -2.0, t)
	})

}

func testInterpretter(expr IExpr, exptdValue float64, t *testing.T) {
	visitor := &Interpreter{}
	numConst := visitor.Visit(expr)
	fmt.Printf("Computer value : %v\n", numConst.Value)
	if exptdValue != numConst.Value {
		t.Error("Interpretter returned a wrong value")
	}
}
