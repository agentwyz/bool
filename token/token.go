package token


type Token interface {
	GetType() Type
	GetText() string
}