# review loop in golang
<br>

### syntax format
_formal syntax in for loop:_
```
    for i:=0;i<=10;i++{
        fmt.Printf("",i)
    }
```
_also we can remove initialisation and post are omitted_
```
    i:=0
    for ;i<=10;i++{
        fmt.Printf("",i)
    }
```

_and we can remove ; in line no .14_
``` 
    i:=0
    for i<=10{
        fmt.Printf("",i)
    }
```
_create a simple calculator_
``` 
package main

import (  
    "fmt"
)

func main() {  
    for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
        fmt.Printf("%d * %d = %d\n", no, i, no*i)
    }

}
```


_infinite loop_
<br>
syntax:
``` 
for {
}
```


<br>
this is a example for infinite loop
<br>


```
package main

import "fmt"

func main() {
	for {
		fmt.Println("hello world")

	}

}
```
# switch statement
==>conditional statement
1)Duplicate cases are not allowed
2)Multiple expressions in case
3)Expressionless switch==>The expression in a switch is optional and it can be omitted


