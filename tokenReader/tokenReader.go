package tokenReader

import "yang/token"

type TokenReader interface {
	peek() token.Type
	read() token.
}