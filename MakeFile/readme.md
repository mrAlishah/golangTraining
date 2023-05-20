#The Makefile
[ref](https://makefiletutorial.com/#pattern-rules)
When you run make, it looks for a file named `Makefile`, or `makefile` in the same directory. The name Makefile is suggested so that it appears near other important files such as README.

You can name your Makefile anything, but then you have to explicitly tell make which file to read:
```makefile
make -f some_other_makefile
```

The Makefile should consist of one or more rules. Each rule describes a goal or a step in your build process, the prerequisites for that step, and recipes for how to execute it.

The format for each rule is as follows:
```shell
target1 [target2 ...]: [pre-req1 pre-req2 pre-req3 ...]
    [recipes
    ...]
```

* The targets are file names, separated by **spaces**. Typically, there is only one per rule.
* The recipes or commands are a series of steps typically used to make the target(s). These need to start with a tab character, not spaces.
* The prerequisites are also file names, separated by **spaces**. These files need to exist before the commands for the target are run. These are also called dependencies

run `make build` to build main.go for first time,the first target is run. In this case, there's target (`build`). The first time you run this, `main` will be created. The second time, you'll see `make: 'main' is up to date.` That's because the main file already exists. But there's a problem: if we modify main.go and then run make, nothing gets recompiled.look at (01-simple)[./01-simple]
```shell
build: main.go
	 go build -o ./main main.go
```

* This has a prerequisite of `main.go`
* Make decides if it should run the build target. It will only run if `main` doesn't exist, or `main.go` is newer than `main`
* This last step is critical, and is the essence of make. What it's attempting to do is decide if the prerequisites of `main` have changed since `main` was last compiled. That is, if `main.go` is modified, running make should recompile the file. And conversely, if `main.go` has not changed, then it should not be recompiled.

To make this happen, it uses the filesystem timestamps as a proxy to determine if something has changed. This is a reasonable heuristic, because file timestamps typically will only change if the files are modified. But it's important to realize that this isn't always the case. You could, for example, modify a file, and then change the modified timestamp of that file to something old. If you did, Make would incorrectly guess that the file hadn't changed and thus could be ignored.

### More quick examples
```shell
01: 02
	echo "01-start"
02: 03
	echo "02-second"
03:
	echo "03-end"
```
**output:** `make 01`
```
echo "03-end"
03-end
echo "02-second"
02-second
echo "01-start"
01-start

```
* Make selects the target `01`, because the first target is the default target
* `01` requires `02`, so make searches for the `02` target
* `02` requires `03`, so make searches for the `03` target
* `03` has no dependencies, so the echo command is run `echo "03-end"`
* The `echo "02-second"` command is then run, because all of the `03` dependencies are finished
* The top `echo "01-start"` command is run, because all the dependencies are finished

### Targets
#### The all target
Making multiple targets and you want all of them to run? Make an all target. Since this is the first rule listed, it will run by default if make is called without specifying a target.
```shell
all: one two three

one:
	@echo "one"
two:
	@echo "two"
three:
	@echo "three"
```
output:
```
one
two
three
```
#### Multiple targets
When there are multiple targets for a rule, the commands will be run for each target. `$@` is an automatic variable that contains the target name.
```shell
all: f1 f2

f1 f2:
	echo $@
# Equivalent to:
# f1:
#	 echo f1
# f2:
#	 echo f2
```

#### Make clean
```shell
some_file: 
	touch some_file

clean:
	rm -f some_file
```
clean is often used as a target that removes the output of other targets, but it is not a special word in Make. You can run make and make clean on this to create and delete some_file.
<br/>
Note that clean is doing two new things here:
* It's a target that is not first (the default), and not a prerequisite. That means it'll never run unless you explicitly call make clean
* It's not intended to be a filename. If you happen to have a file named clean, this target won't run, which is not what we want. See `.PHONY` later in this tutorial on how to fix this

### Variables
Variables can only be strings. You'll typically want to use `:=`, but `=` also works.<br/>
Reference variables using either `${}` or `$()`
```shell
variables := var1 var2
var:
	echo "Look at this variable: " $(variables)
```
Output:
```
echo "Look at this variable: " var1 var2
Look at this variable:  var1 var2
```

another sample:
```shell
files := file1 file2
some_file: $(files)  # we need prerequisite for file1: and file2: to implement
	echo "Look at this variable: " $(files)
	touch some_file

file1:
	touch file1
file2:
	touch file2

clean:
	rm -f file1 file2 some_file
```

and else:
```shell
a := one two # a is set to the string "one two"
b := 'one two' # Not recommended. b is set to the string "'one two'"
print:
	printf '$a'
	@printf $b
```

and note:
```shell
x := dude

all:
	echo $(x)
	echo ${x}

	# Bad practice, but works
	echo $x 
```

#### * Wildcard
Both * and % are called wildcards in Make, but they mean entirely different things. * searches your filesystem for matching filenames. I suggest that you always wrap it in the wildcard function, because otherwise you may fall into a common pitfall described below.
```shell
# Print out file information about every .go file
printfiles: $(wildcard *.go)
	ls -la  $?
```
may be used in the target, prerequisites, or in the wildcard function.
Danger: * may not be directly used in a variable definitions
Danger: When * matches no files, it is left as it is (unless run in the wildcard function)
```shell
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
```

#### Automatic Variables
```shell
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

```

### Read prompt variable
```shell
USERNAME ?= $(shell bash -c 'read -p "Username: " username; echo $$username')
PASSWORD ?= $(shell bash -c 'read -s -p "Password: " pwd; echo $$pwd')

talk:
	#@clear
	@echo "Username › $(USERNAME)"
	@echo "Password › $(PASSWORD)"
```