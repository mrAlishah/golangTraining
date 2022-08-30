# Methods in Go

A method is just a function with a special receiver type between the func keyword and the method name. <br>
The receiver can either be a struct type or non-struct type.
syntax:
```
func (t Type) methodName(parameter list) {  
}

```
# Value receivers in methods vs Value arguments in functions
When a function has a value argument, it will accept only a value argument.

When a method has a value receiver, it will accept both pointer and value receivers.
## Methods with non-struct receivers

So far we have defined methods only on struct types. It is also possible to define methods on non-struct types, but there is a catch


To define a method on a type, the definition of the receiver type and the definition of the method should be present in the same package.
