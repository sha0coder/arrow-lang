package main

import "strings"

const (
	TokStr          = 1
	TokArr          = 2
	TokNum          = 3
	TokVar          = 4
	TokIf           = 5
	TokElse         = 6
	TokReg          = 6
	TokURL          = 8
	TokFile         = 9
	TokGlobal       = 10
	TokFunc         = 11
	TokEof          = 12
	TokIter         = 13
	TokExec         = 14
	TokOper         = 15
	TokAsyncIter    = 16
	TokEndAsyncIter = 17
	TokFuncDef      = 18
)

var Operators = []string{"==", "!=", "=~", "!~", "<", "<=", ">", ">=", "&&", "||", "in"}

func IsComparator(comp string) bool {
	for _, op := range Operators {
		if strings.HasPrefix(comp, op) {
			return true
		}
	}
	return false
}

func GetOperatorAtTheBeginning(op string) string {
	for _, o := range Operators {
		if strings.HasPrefix(op, o) {
			return o
		}
	}
	return ""
}

func GetTokType(tok string) int {
	if len(tok) == 0 {
		return TokEof
	}

	if strings.HasSuffix(tok, ":=>") {
		return TokAsyncIter
	}

	if strings.HasSuffix(tok, "<=:") {
		return TokEndAsyncIter
	}

	if strings.HasSuffix(tok, "=>") {
		return TokIter
	}

	b := tok[0]
	switch b {
	case '[':
		{
			return TokIf
		}
	case '"', '\'':
		{
			return TokStr
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		{
			return TokNum
		}
	case '$':
		{
			return TokVar
		}
	case '@':
		{
			return TokArr
		}
	case '%':
		{
			return TokGlobal
		}
	case '~':
		{
			return TokReg
		}
	case '.', '/':
		{
			return TokFile
		}
	case '|':
		{
			return TokElse
		}
	case '!':
		{
			return TokExec
		}
	case '(':
		{
			return TokOper
		}
	case '*':
		{
			return TokFuncDef
		}

	}
	if len(tok) > 4 && tok[0:4] == "http" {
		return TokURL
	}

	return TokFunc
}
