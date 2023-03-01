### reflect.TypeOf(var)
You can use reflection to get the type of a variable var with the function call `varType := reflect.TypeOf(var)`. This returns a variable of type reflect.Type, which has methods with all sorts of information about the type that defines the variable that was passed in.

### reflect.TypeOf(var).Name() & reflect.TypeOf(var).String()
The first method we’ll look at is `Name()`. This returns, not surprisingly, the name of the type. Some types, like a slice or a pointer, don’t have names and this method returns an empty string.The first method we’ll look at is Name(). This returns, not surprisingly, the name of the type. Some types, like a slice or a pointer, don’t have names and this method returns an empty string.

### reflect.TypeOf(var).Kind()
The next method, and in my opinion the first really useful one, is `Kind()`. The kind is what the type is made of — a slice, a map, a pointer, a struct, an interface, a string, an array, a function, an int or some other primitive type.

### Kind() vs Name()
The difference between the kind and the type can be tricky to understand, but think of it this way. If you define a struct named Foo, the kind is struct and the type is Foo.

#### Warning
One thing to be aware of when using reflection: everything in the reflect package assumes that you know what you are doing and many of the function and method calls will panic if used incorrectly. For example, if you call a method on reflect.Type that’s associated with a different kind of type than the current one, your code will panic. Always remember to use the kind of your reflected type to know which methods will work and which ones will panic.

### reflect.TypeOf(&var).Elem()
If your variable is a pointer, map, slice, channel, or array, you can find out the contained type by using varType.Elem().

### reflect.ValueOf(var) vs reflect.ValueOf(&var)
you can also use reflection to read, set, or create values. First you need to use `refVal := reflect.ValueOf(var)` to create a `reflect.Value` instance for your variable. If you want to be able to use reflection to modify the value, you have to get a pointer to the variable with `refPtrVal := reflect.ValueOf(&var)`; if you don’t, you can read the value using reflection, but you can’t modify it.

### reflect.ValueOf(&var).Set(newRefVal)
If you want to modify a value, remember it has to be a pointer, and you have to dereference the pointer first. You use `refPtrVal.Elem().Set(newRefVal)` to make the change, and the value passed into `Set()` has to be a `reflect.Value` too.

### reflect.New(varType)
If you want to create a new value, you can do so with the function call `newPtrVal := reflect.New(varType)`, passing in a `reflect.Type`. This returns a pointer value that you can then modify. using `Elem().Set()`as described above.