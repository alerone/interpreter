package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

// It gets the value of an identifier stored in the enclosed environment,
//
// if the variable can't be found on the enclosed environment, it searches for the outer one.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// This is needed for functions because they hold their own environment so if the user calls
//
// for an identifier that's in an outer environment without this enclosed characteristics,
//
// the outer variable would be changed with the new value when calling a function.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}
