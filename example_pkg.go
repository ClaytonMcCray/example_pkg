package example_pkg

type ConcreteStruct struct {
	valueOne int
	valueTwo string
}

//testbuilder: methodValueOne on_type *ConcreteStruct args: int
func (c *ConcreteStruct) methodValueOne(x int) {
	c.valueOne = x
}

func (c *ConcreteStruct) methodValueTwo(x string) {
	c.valueTwo = x
}

type InterfaceOfStruct interface {
	methodValueOne(x int)
	methodValueTwo(x string)
}

//testbuilder: doSomething args: *ConcreteStruct, int, string
func doSomething(t InterfaceOfStruct, x int, y string) {
	t.methodValueOne(x)
	t.methodValueTwo(y)
}
