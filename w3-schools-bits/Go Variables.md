## what is variable?
is a named storage that holds a value which can change during program execution.
Variables are declared using the var keyword (or with := for shorthand)

## types of variables

- int – stores whole numbers, like 123 or -123
- float32 – stores numbers with decimals, such as 19.99 or -19.99
- string – stores text, for example "Hello World". Strings are enclosed in double quotes
- bool – stores a logical value, either true or false

### Declaring variables

1. Using var keyword

var name string = "Alice"
var age int = 25
using this method we use the var keyword followed by the variable name and type

2. Using shorthand :=
name := "Bob"


#### Multiple variable declaration

this is a way of declaring multiple variables in one line instead of writing separate lines for each variable. e.g. var a, b, c, d int = 1, 3, 5, 7 If you use the type keyword, it is only possible to declare one type of variable per line. If the type keyword is not specified, you can declare different types of variables on the same line

#### Variable declaration in a block

This means declaring multiple variables together inside a pair of parentheses using the var keyword. e.g. var ( a int b int = 1 c string = "hello" )

#### variable naming rules

This are the guildlines that determines how someone can correctly name variables in Go. Variable name containletters, numbers,and undesrscores It cannot start with a number it cannot use Go keywords(e.g func, var, package) Names are case-sensitive (age and Age are different) These rules ensure your variable names are valid and understood by the Go compiler.

### Multi word variable

This are variables made up of more than one word, written in a readable format. In Go, they are usually written using camelCase, where the first word is lowercase and each new word starts with a capital letter. e.g. UserName := "naeema"

