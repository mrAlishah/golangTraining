
# Array
<br>

Arrays are fixed-size collections of same-type items
<br>
Arrays are accessed using an array index
<br>
It is an error to use an index outside the bounds of an array
<br>
Array elements can be optionally set during array creation
<br>
Elements not manually assigned a value will have a default Use the len() function to iterate arrays in a for loop
<br>




# slice

write a slice with this command=>[]Type
<br>Slices do not own any data on their own. They are just references to existing arrays.
<br> The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
<br>syntax capacity==>
```
fmt.Printf("",len(),cap())
```
### creating a slice using make

<br>func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity
`make([]int,5,5)`=>out put:[0 0 0 0 0]
### Appending to a slice
vim:arrays are restricted to fixed length and their length cannot be increased.
very good point:when new elements are appended to the slice, a new array is created.
in ex number 45 The capacity of the new slice is now twice that of the old slice
**The zero value of a slice type is nil.**
<br>
<br>
<br>
`func append(s []T, x ...T) []T`
### Passing a slice to a function

```
type slice struct {  
    Length        int
    Capacity      int
    ZerothElement *byte
}
```

slice like array can have multiple dimensions.
<br>
syntax:
`[][]string `


# learn about **Variadic Functions**
<br>
A variadic function is a function that accepts a variable number of arguments
<br>
Only the last parameter of a function can be variadic
<br>
syntax:
<br>
```
func funcname(var-name type,var name ...type){
}
```


very imp:It is also possible to pass zero arguments to a variadic function.

The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter
