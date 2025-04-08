package geno

type Token interface {
	Parse(*Parser) (bool, ParseHandler)
}

type PositionedToken struct {
	Token     Token
	CursorPos int // The cursor position immedtely to the left of the token in the source file
}

func NewPositionedToken(t Token, cursorPos int) PositionedToken {
	return PositionedToken{
		Token:     t,
		CursorPos: cursorPos,
	}
}
