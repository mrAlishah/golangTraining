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

### 05) Making Without Make
In addition to making instances of built-in and user-defined types, you can also use reflection to make instances that normally require the make function. You can make a slice, map, or channel using the `reflect.MakeSlice`, `reflect.MakeMap`, and `reflect.MakeChan` functions. In all cases, you supply a `reflect.Type` and get back a `reflect.Value` that you can manipulate with reflection, or that you can assign back to a standard variable.

### 06) Making Functions
Reflection doesn’t just let you make new places to store data. You can use reflection to make new functions using the `reflect.MakeFunc` function. This function expects the `reflect.Type` for the function that we want to make and a closure whose input parameters are of type []reflect.Value and whose output parameters are also of type []reflect.Value

### 07) I Want a New Struct
There’s one more thing that you can make using reflection in Go. You can make brand-new structs at runtime by passing a slice of reflect.StructField instances to the `reflect.StructOf` function. This one is a bit weird; we are making a new type, but we don’t have a name for it, so you can’t really turn it back into a “normal” variable. You can create a new instance and use Interface() to put the value into a variable of type interface{}, but if you want to set any values on it, you need to use reflection