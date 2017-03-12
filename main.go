package main

import (
	"bufio"
	"fmt"
	"os"
	"slango/slang"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Exp> ")
		expStr, _ := reader.ReadString('\n')
		if 0 == strings.Compare("exit\n", expStr) {
			fmt.Println("Good Bye!")
			os.Exit(0)
		}
		expr := slang.ParseFromStream([]byte(expStr))
		interpreter := &slang.Interpreter{}
		numConst := interpreter.Visit(expr)
		fmt.Print(numConst.Value)
		fmt.Println()
	}
}
