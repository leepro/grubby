package builtins

import (
	"errors"
	"fmt"
)

type fixnumClass struct {
	valueStub
	classStub

	instanceMethods []Method
}

func NewFixnumClass(provider ClassProvider) Class {
	class := &fixnumClass{}
	class.initialize()
	class.class = provider.ClassWithName("Class")
	class.superClass = provider.ClassWithName("Integer")

	return class
}

func (class *fixnumClass) AddInstanceMethod(method Method) {
	class.instanceMethods = append(class.instanceMethods, method)
}

func (c *fixnumClass) String() string {
	return "Fixnum"
}

func (c *fixnumClass) Name() string {
	return "Fixnum"
}

func (c *fixnumClass) New(provider ClassProvider, args ...Value) (Value, error) {
	return nil, errors.New("undefined method 'new' for Fixnum:Class")
}

type fixnumInstance struct {
	value int
	valueStub
}

func NewFixnum(val int, provider ClassProvider) Value {
	i := &fixnumInstance{value: val}
	i.class = provider.ClassWithName("Fixnum")
	i.initialize()
	return i
}

func (fixnumInstance *fixnumInstance) Value() int {
	return fixnumInstance.value
}

func (fixnumInstance *fixnumInstance) String() string {
	return fmt.Sprintf("%d", fixnumInstance.value)
}