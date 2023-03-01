# Type assertions vs. type conversions

## What is type assertion?
Type assertion (as the name implies) is used to assert the type of a given variable. In Go, this is done by checking the underlying type of an empty interface variable.

## What is type conversion?
Type conversion is the process of changing a variable from one type to another specified type. For example, we can convert an int value to a float64.

### Sample 02-assertion.go 
A type assertion provides access to an interface value's underlying concrete value.
```go
t := i.(T)
```

This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
```go
t, ok := i.(T)
```

If `i` holds a `T`, then t will be the underlying value and `ok` will be true.

If not, `ok` will be false and `t` will be the zero value of type `T`, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.