# makefile --------------------------------------------
hello:
	echo "Hello, World"
	echo "This line will always print, because the file hello does not exist."

# makefile --------------------------------------------
build: main.go
	 go build -o ./main main.go

# More quick examples ---------------------------------
01: 02
	echo "01-start"
02: 03
	echo "02-second"
03:
	echo "03-end"

# The all target --------------------------------------
all: one two three

one:
	@echo "one"
two:
	@echo "two"
three:
	@echo "three"

# Multiple targets --------------------------------------
all1: f1 f2

f1 f2:
	echo $@
# Equivalent to:
# f1:
#	 echo f1
# f2:
#	 echo f2

# Make clean -----------------------------------------
some_file:
	touch some_file
clean:
	rm -f some_file

# Variables ------------------------------------------
variables := var1 var2
var:
	echo "Look at this variable: " $(variables)

# Variables ------------------------------------------
a := one two # a is set to the string "one two"
b := 'one two' # Not recommended. b is set to the string "'one two'"
print:
	@printf a:'$a'
	@printf b:$b

# Variables Wildcard ---------------------------------
# Print out file information about every .go file
printfiles: $(wildcard *.go)
	ls -la  $?

# Variables Wildcard ---------------------------------
thing_wrong := *.go # Don't do this! '*' will not get expanded
thing_right := $(wildcard *.go)

cards: cardone cardtwo cardthree cardfour

# Fails, because $(thing_wrong) is the string "*.go"
cardone: $(thing_wrong)
	echo "cardone"
# Stays as *.go if there are no files that match this pattern :(
cardtwo: *.go
	echo "cardtwo"
# Works as you would expect! In this case, it does nothing.
cardthree: $(thing_right)
	echo "cardthree"
# Same as rule three
cardfour: $(wildcard *.go)
	echo "cardfour"


# Automatic Variables ---------------------------------
Auto: Auto1 Auto2
	# Outputs "Auto", since this is the target name
	echo $@

	# Outputs all prerequisites newer than the target
	echo $?

	# Outputs all prerequisites
	echo $^

	echo "hey"

Auto1:
	echo "one"

Auto2:
	echo "two"
