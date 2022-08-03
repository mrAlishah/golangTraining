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
    

 




