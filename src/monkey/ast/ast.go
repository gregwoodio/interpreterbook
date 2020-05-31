package ast

import (
	"bytes"
	"monkey/token"
)

// Node is a node in the AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is an AST node representing a statement
type Statement interface {
	Node
	statementNode()
}

// Expression is an AST node representing an expression
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

// String representation of Program
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String representation of LetStatement
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier is an AST node representing an identifier
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// String representation of Identifier
func (i *Identifier) String() string {
	return i.Value
}

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

// String representation of ReturnStatement
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is an AST node representing an expression statement.
// This exists so we can add an expression to the list of statements in our // program. ie. `5 + 5` on a line by itself is valid in Monkey
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral for ExpressionStatement token
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// IntegerLiteral is an expression representing an integer
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral for IntegerLiteral
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression is an expression before another expression, such as -1 or !true
type PrefixExpression struct {
	Token    token.Token //the prefix token, eg !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral for PrefixExpression
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression is an expression between two expresssions, such as '5 + 5' or '10 / 2'
type InfixExpression struct {
	Token    token.Token // the operator token, eg +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral for InfixExpression
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean represents a boolean value
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral for Boolean
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}
