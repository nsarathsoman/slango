package slang

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	// test case 1
	t.Run("case1", func(t *testing.T) {
		slangTestFile := "/tmp/num1.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("1234"), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if 1234 != numericConst.Value {
			t.Errorf("Parse error expected 1234 got %v", numericConst.Value)
		}
	})

	//test case 2
	t.Run("case2", func(t *testing.T) {
		slangTestFile := "/tmp/exp2.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("1234 + 1"), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if 1235 != numericConst.Value {
			t.Errorf("Parse error expected 1235 got %v", numericConst.Value)
		}
	})

	//test case 3
	t.Run("case3", func(t *testing.T) {
		slangTestFile := "/tmp/exp3.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("1234 * 2"), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if 2468 != numericConst.Value {
			t.Errorf("Parse error expected 2468 got %v", numericConst.Value)
		}
	})

	//test case 4
	t.Run("case4", func(t *testing.T) {
		slangTestFile := "/tmp/exp4.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("20+3*4-1"), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if 31 != numericConst.Value {
			t.Errorf("Parse error expected 31 got %v", numericConst.Value)
		}
	})

	//test case 5
	t.Run("case5", func(t *testing.T) {
		slangTestFile := "/tmp/exp5.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("20+3/-4-1   "), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if 18.25 != numericConst.Value {
			t.Errorf("Parse error expected 18.25 got %v", numericConst.Value)
		}
	})

	//test case 6
	t.Run("case6", func(t *testing.T) {
		slangTestFile := "/tmp/exp6.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("20+3/+4-1   "), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if 19.75 != numericConst.Value {
			t.Errorf("Parse error expected 19.75 got %v", numericConst.Value)
		}
	})

	//test case 7
	t.Run("case7", func(t *testing.T) {
		slangTestFile := "/tmp/exp7.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("2 + (10 / -5) * (-1)   "), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if 4 != numericConst.Value {
			t.Errorf("Parse error expected 4 got %v", numericConst.Value)
		}
	})

	//test case 8
	t.Run("case8", func(t *testing.T) {
		slangTestFile := "/tmp/exp8.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("(10 / -5) * (-1) + 2  "), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if 4 != numericConst.Value {
			t.Errorf("Parse error expected 4 got %v", numericConst.Value)
		}
	})

	//test case 9
	t.Run("case9", func(t *testing.T) {
		slangTestFile := "/tmp/exp9.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("56/43*5+45-67"), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if -15.488372093 != numericConst.Value {
			t.Errorf("Parse error expected −15.488372093 got %v", numericConst.Value)
		}
	})

	//test case 10
	t.Run("case10", func(t *testing.T) {
		slangTestFile := "/tmp/exp10.slang"
		err := ioutil.WriteFile(slangTestFile, []byte("23/90*6+65-54"), 0644)
		checkErr(err)
		expr := parse(slangTestFile)
		interpreter := &Interpreter{}
		numericConst := interpreter.Visit(expr)
		if 12.533333333 != numericConst.Value {
			t.Errorf("Parse error expected 12.533333333 got %v", numericConst.Value)
		}
	})

}
