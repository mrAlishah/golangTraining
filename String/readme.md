# **string**
_A string is a slice of bytes in Go_<br>
_Strings in Go are Unicode compliant and are UTF-8 Encoded_<br>
_Since a string is a slice of bytes, it's possible to access each byte of a string._<br>
_The == operator is used to compare two strings for equality_<br>
_In UTF-8 encoding a code point can occupy more than 1 byte so we use rune_<br>
_review rune:A rune is a builtin type in Go and it's the alias of int32_<br>
_The RuneCountInString(s string) (n int) function of the utf8 package can be used to find the length of the string
This method takes a string as an argument and returns the number of runes in it_<br>
_Strings are immutable in Go. Once a string is created it's not possible to change it_<br>
_To workaround this string immutability, strings are converted to a slice of runes.
Then that slice is mutated with whatever changes are needed and converted back to a new string_<br>
## Trim
1. Trim: This function is used to trim the string all the leading and trailing Unicode code points which are specified in this function.
2. TrimLeft: This function is used to trim the left-hand side(specified in the function) Unicode code points of the string.
3. TrimRight: This function is used to trim the right-hand side(specified in the function) Unicode code points of the string.
4. TrimSpace: This function is used to trim all the leading and trailing white space from the specified string.
5. TrimSuffix: This method is used to trim the trailing suffix string from the given string. If the given string does not contain the specified suffix string, then this function returns the original string without any change.
6. TrimPrefix: This method is used to trim the leading prefix string from the given string. If the given string does not contain the specified prefix string, then this function returns the original string without any change.



 




