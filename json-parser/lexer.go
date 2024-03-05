package main

import (
	"errors"
	"fmt"
)

type Lexer struct {
	Input    []rune
	Position int
}

func NewLexer(input string) Lexer {
	return Lexer{Input: []rune(input)}
}

func (l *Lexer) skipWhitespaces() {
	for {
		switch l.Input[l.Position] {
		case ' ', '\b', '\t', '\n', '\r':
			l.Position++
			continue
		}
		break
	}
}

func (l *Lexer) readString(start int) (string, error) {
	l.Position++
	for l.Position < len(l.Input) && l.Input[l.Position] != '"' {
		l.Position++
	}
	if l.Position < len(l.Input) {
		return string(l.Input[start+1 : l.Position]), nil
	}
	return "", errors.New(`invalid json value, expect ", but not found`)
}

func (l *Lexer) readValue(start int) (string, Type, error) {
	switch l.Input[start] {
	case 't':
		if start+4 < len(l.Input) {
			str := string(l.Input[start : start+4])
			if str == "true" {
				l.Position += 5
				return "true", True, nil
			}
		}
		return "", Invalid, errors.New("invalid json value, expect true")
	case 'f':
		if start+5 < len(l.Input) {
			str := string(l.Input[start : start+5])
			if str == "false" {
				l.Position += 6
				return "false", False, nil
			}
		}
		return "", Invalid, errors.New("invalid json value, expect false")
	default:
		if start+4 < len(l.Input) {
			str := string(l.Input[start : start+4])
			if str == "null" {
				l.Position += 5
				return "null", Null, nil
			}
		}
		return "", Invalid, errors.New("invalid json value, expect null")
	}
}

func (l *Lexer) readNumber(start int) (string, error) {
	// TODO: frac, exp
	if l.Input[l.Position] == '-' && l.Input[l.Position+1] == '.' {
		return "", errors.New("invalid json value, expect number")
	}
	l.Position++
	for {
		if l.Position >= len(l.Input) {
			return "", errors.New("invalid json value, expect number")
		}
		c := l.Input[l.Position]
		if ('0' <= c && c <= '9') || c == '.' {
			l.Position++
		} else {
			break
		}
	}
	return string(l.Input[start:l.Position]), nil
}

func (l *Lexer) Scan() (Token, error) {
	if l.Position == len(l.Input) {
		return Token{Type: EOF, Literal: "", Start: l.Position, End: l.Position}, nil
	}

	l.skipWhitespaces()

	switch c := l.Input[l.Position]; c {
	case '{':
		t := Token{Type: LeftBrace, Literal: "{", Start: l.Position, End: l.Position}
		l.Position++
		return t, nil
	case '}':
		t := Token{Type: RightBrace, Literal: "}", Start: l.Position, End: l.Position}
		l.Position++
		return t, nil
	case '[':
		t := Token{Type: LeftBracket, Literal: "[", Start: l.Position, End: l.Position}
		l.Position++
		return t, nil
	case ']':
		t := Token{Type: RightBracket, Literal: "]", Start: l.Position, End: l.Position}
		l.Position++
		return t, nil
	case ':':
		t := Token{Type: Colon, Literal: ":", Start: l.Position, End: l.Position}
		l.Position++
		return t, nil
	case ',':
		t := Token{Type: Comma, Literal: ",", Start: l.Position, End: l.Position}
		l.Position++
		return t, nil
	case '"':
		start := l.Position
		str, err := l.readString(start)
		if err != nil {
			return Token{}, err
		}
		t := Token{Type: String, Literal: str, Start: start + 1, End: l.Position - 1}
		l.Position++
		return t, nil
	case 't', 'f', 'n':
		start := l.Position
		str, t, err := l.readValue(start)
		if err != nil {
			return Token{}, err
		}
		return Token{Type: t, Literal: str, Start: start, End: l.Position}, nil
	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		start := l.Position
		str, err := l.readNumber(start)
		if err != nil {
			return Token{}, err
		}
		t := Token{Type: Number, Literal: str, Start: start, End: l.Position - 1}
		return t, nil
	}
	return Token{}, errors.New(fmt.Sprintf("invalid json value, position=%d, char=%s", l.Position, string(l.Input[l.Position])))
}
