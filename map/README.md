# How to create a map?

syn:
``` 
  make(map[type of key]type of value)
```
# Adding items to a map
syn:
```
f:=map[string]int{
    "saman":2000
    "saeid":1500
    "shima":3000
    "parto":2500
    }
||
var f map[string]int{
    "salad":55
    "chicken":45
    "rice":32
   } 
||
f:=make(map[strng]int)
    f["saman"]=2000
    f["saeid"]=1500
    
    

```
# Retrieving value for a key from a map
```
f:=map[string]int{
    "saman":2000
    "saeid":1500
    "shima":3000
    "parto":2500
    }
    j:="shima"
    d:=f[j]// so here we have retriev shima
    fmt.println(d)
    

```

# Checking if a key exists
syn:
`value, ok := map[key]  
`
# Iterate over all elements in a map
use range
for key, value := range employeeSalary {<br>
}

# Deleting items from a map
delete(map, key) is the syntax to delete key from a map. The delete function does not return any value.

# Length of the map
using the len function

# very important point 
**Similar to slices, maps are reference types.**<br>
Similar is the case when maps are passed as parameters to functions.<br>
When any change is made to the map inside the function, it will be visible to the caller also.
look at example number 87   
 # Maps equality
Maps can't be compared using the == operator. The == can be only used to check if a map is nil.
``` 
package main

func main() {  
    map1 := map[string]int{
        "one": 1,
        "two": 2,
    }

    map2 := map1

    if map1 == map2 {
    }
}
```
invalid operation: map1 == map2 (map can only be compared to nil)  