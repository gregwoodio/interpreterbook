package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is the root of our AST, and it contains a slice of statements
type Program struct {
	Statements []Statement
}

// TokenLiteral for Program node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement is an AST node representing a let statement
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral for LetStatement token
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier is an AST node representing an identifier
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral for Identifier token
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ReturnStatement is an AST node representing a return statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral for ReturnStatement token
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
