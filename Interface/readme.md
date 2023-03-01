# interface

In Go, an interface is a set of method signatures.<br>
When a type provides definition for all the methods in the interface, it is said to implement the interface
<br><br>good example for understanding interface:
<br>For example WashingMachine can be an interface with method signatures Cleaning() and Drying(). Any type which provides definition for Cleaning() and Drying() methods is said to implement the WashingMachine interface.

<br>
 An interface can be thought of as being represented internally by a tuple (type, value). type is the underlying concrete type of the interface and value holds the value of the concrete type.

# Empty interface
An interface that has zero methods is called an empty interface. It is represented as interface{}
# type assertion

Type assertion is used to extract the underlying value of the interface.
### What is type assertion?
Type assertion (as the name implies) is used to assert the type of a given variable. In Go, this is done by checking the underlying type of an empty interface variable.

### What is type conversion?
Type conversion is the process of changing a variable from one type to another specified type. For example, we can convert an int value to a float64.
i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.
# type switch
A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements.
It is similar to switch case. 
The only difference being the cases specify types and not values as in normal switch.
<br>
It is also possible to compare a type to an interface. If we have a type and if that type implements an interface, it is possible to compare this type with the interface it implements.<br>
<br>

The reason is that it is legal to call a pointer-valued method on anything that is already a pointer or whose address can be taken. The concrete value stored in an interface is not addressable and hence it is not possible for the compiler to automatically take the address.
<br>
<br>
<br>
<br>
A type can implement more than one interface
<br>
<br>

The zero value of a interface is nil. A nil interface has both its underlying value and as well as concrete type as nil.