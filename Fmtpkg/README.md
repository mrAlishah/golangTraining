# fmt tutorial

<br>

%d=>for int=>decimal
<br>
%s=>for string

%g=>for float

%e & %E =>for scientific values like constant

%b=>convert base 10 to base 2

%o=>convert base 10 to base 8=>octal

%O=>convert base 10 to base 8 with 0o prefix

%%=>no value just type %
%x=>convert base 10 to base 16=> upper-case letters for A-F
%X=>convert base 10 to base 16=>hexadecimal
%#X=>base 16 with leading 0x
%q=>string format with quoted
%c=>for rune(c => character)
%f=>float like %g but with width & precision
Ex:
%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0


٪v =>for struct
%#v=>for object & struct
%+v=>adds field names
%t=>true or false
%p=>to print a representation of a pointer
















Sprintf=>format a string without printing it
