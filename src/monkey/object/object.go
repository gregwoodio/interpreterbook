package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
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
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
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

// HashKey for Integer
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
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

// HashKey for Boolean
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
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

// HashKey for String
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
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

// Array represents an array
type Array struct {
	Elements []Object
}

// Type for Array
func (ao *Array) Type() ObjectType {
	return ARRAY_OBJ
}

// Inspect for Array
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashKey represents a hashed value
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashPair represents a Key Value pair of Objects
type HashPair struct {
	Key   Object
	Value Object
}

// Hash is a dictionary between HashKey and HashPair
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type for Hash
func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

// Inspect for Hash
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Hashable interface is for Objects that can be Hashed, like Strings, Integers and Booleans
type Hashable interface {
	HashKey() HashKey
}
