package main

import (
	"fmt"
	"yang/token"
	"yang/dfastate"
)


type SimpleToken struct {
	Text string
	Type token.Type
}

func main() {
	var code = "+"
	var ans  dfastate.State = tokenize(code)
	fmt.Println(ans)

	fmt.Println(stoken.Text)
	fmt.Println(stoken.Type)

	code = ">"
	ans = tokenize(code)
	fmt.Println(ans)

	fmt.Println(stoken.Text)
	fmt.Println(stoken.Type)

	code = ">="
	ans = tokenize(code)
	fmt.Println(ans)

	fmt.Println(stoken.Text)
	fmt.Println(stoken.Type)
}

var stoken SimpleToken
var tokenText string

func tokenize(code string) dfastate.State {
	var state dfastate.State = dfastate.Init
	for _, ch := range code {
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
		}

		//需要将最后一个东西送进去
		if len(tokenText) > 0 {
			initToken(byte(ch))
		}
	}

	return state
}

func initToken(ch byte) dfastate.State {
	if len(tokenText) > 0 {
		stoken.Text = tokenText

		tokenText = ""
	}

	var state dfastate.State = dfastate.Init
	if isDigit(ch) {
		stoken.Type = token.INT
		state = dfastate.IntLiteral
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