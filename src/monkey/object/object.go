package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

// ObjectType is the type of the object
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
)

// Object representation used in the evaluator
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer object type
type Integer struct {
	Value int64
}

// Inspect for Integer
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type for Integer
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Boolean object type
type Boolean struct {
	Value bool
}

// Inspect for Boolean
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type for Boolean
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Null object type
type Null struct{}

// Inspect for Null
func (n *Null) Inspect() string {
	return "null"
}

// Type for Null
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// ReturnValue object type
type ReturnValue struct {
	Value Object
}

// Type for ReturnValue
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// Inspect for ReturnValue
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error type
type Error struct {
	Message string
}

// Type for Error
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

// Inspect for Error
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Function object type
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type for Function
func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

// Inspect for Function
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String Object type
type String struct {
	Value string
}

// Type for String
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect for String
func (s *String) Inspect() string {
	return s.Value
}

// BuiltinFunction is the underlying Go function called
type BuiltinFunction func(args ...Object) Object

// Builtin functions wrapper
type Builtin struct {
	Fn BuiltinFunction
}

// Type for Builtin
func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

// Inspect for Builtin
func (b *Builtin) Inspect() string {
	return "builtin function"
}
