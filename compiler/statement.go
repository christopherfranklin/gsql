package compiler

import "strings"

type PrepareResult int8

const (
	SUCCESS PrepareResult = iota
	UNRECOGNIZED_STATEMENT
)

type StatementType int

const (
	INSERT StatementType = iota
	SELECT

	ERROR
)

type Statement struct {
	StatementType StatementType
}

func PrepareStatement(input string) (Statement, PrepareResult) {
	if strings.HasPrefix(input, "insert") {
		return Statement{INSERT}, SUCCESS
	} else if strings.HasPrefix(input, "select") {
		return Statement{SELECT}, SUCCESS
	} else {
		return Statement{ERROR}, UNRECOGNIZED_STATEMENT
	}
}
