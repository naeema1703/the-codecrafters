## OUTPUT 
output functions usually refer to the built-in functions in the fmt package used to print text to the standard output (the console). 

###

print() > fmt.Print(): Prints arguments in their default format. It does not add a new line at the end or spaces between arguments unless they are not strings.
println() Similar to Print(), but it always adds a space between multiple arguments and appends a newline (\n) at the end of the output.
print() it is Used for formatted output. It uses "formatting verbs" (like %v for value or %T for type) to specify exactly how the data should be displayed. 