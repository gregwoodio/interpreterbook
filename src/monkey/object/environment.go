package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{
		store: s,
		outer: nil,
	}
}

// Environment type keeps track of our variables
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get a variable's value
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set a variable's value
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
