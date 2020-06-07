package object

import "fmt"

// ObjectType is the type of the object
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
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
