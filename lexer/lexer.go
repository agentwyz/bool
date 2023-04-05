package main

import (
	"fmt"
	"yang/dfastate"
	"yang/token"
)

type Token interface {
	getType() token.Type
	getText() string
}

type SimpleToken struct {
	Text string
	Type token.Type
}

func (stk SimpleToken) getText() string {
	return stk.Text
}

func (stk SimpleToken) getType() token.Type {
	return stk.Type
}

func main() {
	var code string = "+>>"
	var ans dfastate.State = tokenize(code)
	_ = ans
	code = ">"
	ans = tokenize(code)

	code = ">="
	ans = tokenize(code)

	code = "+"
	fmt.Print("==========\n")
	ans = tokenize(code)
	fmt.Println(ans)

	for i, j := range tokens {
		fmt.Print(i)
		fmt.Print("\t\t")
		fmt.Println(j)
	}

	code = "+>+"
	ans = tokenize(code)
	for i, j := range tokens {
		fmt.Print(i)
		fmt.Print("\t\t")
		fmt.Println(j)
	}
	fmt.Println("============================")

	//重新更新
	tokens = []Token{}
	code = "int"
	ans = tokenize(code)
	for i, j := range tokens {
		fmt.Print(i)
		fmt.Print("\t\t")
		fmt.Println(j)
	}
}

var stoken SimpleToken
var tokenText string
var tokens []Token

func tokenize(code string) dfastate.State {
	var state dfastate.State = dfastate.Init
	tokenText = ""
	stoken = SimpleToken{}
	var ch rune
	for _, ch = range code {
		switch state {
		case dfastate.Init:
			state = initToken(byte(ch))

		case dfastate.Plus:
			state = initToken(byte(ch))

		case dfastate.GT:
			if ch == '=' {
				tokenText += string(ch)
				stoken.Type = token.GE
				state = dfastate.GE
			} else {
				state = initToken(byte(ch))
			}

		case dfastate.GE:
			state = initToken(byte(ch))

		case dfastate.Assign:
			state = initToken(byte(ch))

		case dfastate.EQ:
			state = initToken(byte(ch))

		case dfastate.Id:
			if isAlpha(byte(ch)) {
				tokenText += string(ch)
			} else {
				state = initToken(byte(ch))
			}

		default:
		}
	}

	//这个地方需要非常的小心
	if len(tokenText) > 0 {
		initToken(byte(ch))
	}

	return state
}

func initToken(ch byte) dfastate.State {
	if len(tokenText) > 0 {

		stoken.Text = tokenText

		tokens = append(tokens, stoken)
		tokenText = ""
		stoken = SimpleToken{}
	}

	var state dfastate.State = dfastate.Init
	if isDigit(ch) {
		stoken.Type = token.INT
		state = dfastate.IntLiteral
		tokenText += string(ch)
	} else if isAlpha(ch) {
		stoken.Type = token.IDENT
		state = dfastate.Id
		tokenText += string(ch)
	} else if ch == '+' {
		stoken.Type = token.PLUS
		state = dfastate.Plus
		tokenText += string(ch)
	} else if ch == '>' {
		stoken.Type = token.GT
		state = dfastate.GT
		tokenText += string(ch)
	}

	return state
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func isBlank(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
