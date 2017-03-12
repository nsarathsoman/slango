package slang

import (
	"io/ioutil"
	"testing"
)

func TestReadNum(t *testing.T) {
	//test case 1
	t.Run("case1", func(t *testing.T) {
		err := ioutil.WriteFile("/tmp/num1.slang", []byte("1234"), 0644)
		checkErr(err)
		data, err := ioutil.ReadFile("/tmp/num1.slang")
		checkErr(err)
		lexer := &Lexer{Data: data, DataLen: len(data), Index: 0}
		lexer.readNum()
		if lexer.Num != 1234 {
			t.Error("Parsed number is not the expexted one")
		}
	})

	//test case 2
	t.Run("case2", func(t *testing.T) {
		data := []byte("123.45 ")
		lexer := &Lexer{Data: data, DataLen: len(data), Index: 0}
		lexer.readNum()
		if lexer.Num != 123.45 {
			t.Error("Parsed number is not the expexted one")
		}
	})

	//test case 3
	t.Run("case3", func(t *testing.T) {
		data := []byte("123.4.5")
		lexer := &Lexer{Data: data, DataLen: len(data), Index: 0}
		defer func() {
			if rec := recover(); rec == nil {
				t.Error("readNum() is not pancing for an invalid input")
			}
		}()
		lexer.readNum()
	})

}

func TestEat(t *testing.T) {
	//test case1
	t.Run("case1", func(t *testing.T) {
		data := []byte("123.4")
		lexer := &Lexer{Data: data, DataLen: len(data), Index: 0}
		lexer.eat()
		token := lexer.CurToken
		if 123.4 != lexer.Num {
			t.Errorf("Eat function did not extract the numeric value properly, value %v", lexer.Num)
		}

		if NUM != token {
			t.Errorf("Identified token is not Num, Token : %v", token)
		}
	})

	//test case2
	t.Run("case2", func(t *testing.T) {
		data := []byte("123.4 + 3")
		lexer := &Lexer{Data: data, DataLen: len(data), Index: 0}
		lexer.eat()
		token := lexer.CurToken
		if 123.4 != lexer.Num {
			t.Errorf("Eat function did not extract the numeric value properly, value %v", lexer.Num)
		}

		if NUM != token {
			t.Errorf("Identified token is not Num, Token : %v", token)
		}

		lexer.eat()
		token = lexer.CurToken

		if ADD != token {
			t.Errorf("Identified token is not Add, Token : %v", token)
		}

		lexer.eat()
		token = lexer.CurToken

		if NUM != token {
			t.Errorf("Identified token is not Num, Token : %v", token)
		}

		if 3 != lexer.Num {
			t.Errorf("Eat function did not extract the numeric value properly, value %v", lexer.Num)
		}

		lexer.eat()
		token = lexer.CurToken
		if UNKNOWN != token {
			t.Errorf("Identified token is not UNKNOWN, Token : %v", token)
		}
	})
}

func TestEatParan(t *testing.T) {
	//test case1
	t.Run("case1", func(t *testing.T) {
		data := []byte("(  )")
		lexer := &Lexer{Data: data, DataLen: len(data), Index: 0}
		lexer.eat()
		token := lexer.CurToken
		if OPAR != token {
			t.Errorf("Identified token is not OPAR, Token : %v", token)
		}

		lexer.eat()
		token = lexer.CurToken
		if CPAR != token {
			t.Errorf("Identified token is not CPAR, Token : %v", token)
		}
	})
}

func TestExcept(t *testing.T) {
	//test case1
	t.Run("case1", func(t *testing.T) {
		data := []byte("( *")
		lexer := &Lexer{Data: data, DataLen: len(data), Index: 0}
		lexer.eat()
		if OPAR != lexer.CurToken {
			t.Errorf("Identified token is not OPAR, Token : %v", lexer.CurToken)
		}
		defer func() {
			if rec := recover(); rec == nil {
				t.Error("expect function not panicing when an expectation is not met")
			}
		}()
		lexer.eat()
		lexer.expect(CPAR)
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
