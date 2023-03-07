# dorklang

This is an interpreter, written in Go, for **dorklang**, which is an esoteric programming language devised by me.

## Examples

After cloning the repository to your local machine, try running one of the example programs with the command below:

```
go run .\error.go .\flag.go .\lexeme.go .\main.go .\tree.go "--file=examples/countdown.dork"
```

## Syntax

Below is an overview of all the commands that can be used in **dorklang** source-code files:

| Command | Function |
| :--------: | ------- |
| `+` | Adds `1` to the current value. |
| `++` | Adds `8` to the current value. |
| `%+` | Pops the two topmost values from the current stack, adds one to the other and sets the current value to the result. | 
| `%++` | Pops all of the values from the current stack, adds each of them to the others and sets the current value to the result. | 
| `-` | Subtracts `1` from the current value. |
| `--` | Subtracts `8` from the current value. |
| `%-` | Pops the two topmost values from the current stack, subtracts one from the other and sets the current value to the result. | 
| `%--` | Pops all of the values from the current stack, subtracts each of them from the others and sets the current value to the result. | 
| `/` | Divides the current value by `2`. |
| `//` | Divides the current value by `8`. |
| `%/` | Pops the two topmost values from the current stack, divides one from the other and sets the current value to the result. | 
| `%//` | Pops all of the values from the current stack, divides each of them from the others and sets the current value to the result. | 
| `*` | Multiplies the current value by `2`. |
| `**` | Multiplies the current value by `8`. |
| `%*` | Pops the two topmost values from the current stack, multiplies one with the other and sets the current value to the result. | 
| `%**` | Pops all of the values from the current stack, multiplies each of them with the others and sets the current value to the result. | 
| `^` | Squares the current value (i.e. multiplies it by itself). |
| `^^` | Cubes the current value (i.e. multiplies it by itself twice). |
| `!` | Prints the current value to the screen as a Unicode/ASCII character. |
| `!!` | Prints the current value to the screen as a decimal number. |
| `?` | Waits for a Unicode/ASCII character to be given as input, then sets the current value to its numerical value. |
| `??` | Waits for a decimal number to be given as input, then sets the current value to it. |
| `~` | Sets the current value to `0`. |
| `'` | Sets the current value to the size of a byte (i.e. `8`). |
| `''` | Sets the current value to the size of eight bytes (i.e. `64`). |
| `"` | Sets the current value to the size of a kibibyte (i.e. `8_192`). |
| `""` | Sets the current value to the size of eight kibibytes (i.e. `65_536`). |
| `%'` | Sets the current value to the size of a mebibyte (i.e. `8_388_608`). |
| `%''` | Sets the current value to the size of eight mebibytes (i.e. `67_108_864`). |
| `%"` | Sets the current value to the size of a gibibyte (i.e. `8_589_934_592`). |
| `%""` | Sets the current value to the size of eight gibibytes (i.e. `68_719_476_736`). |
| `\`` | Sets the current value to a random number between `0` and `255`. |
| `\`\`` | Sets the current value to a random number between `0` and the maximum value for an unsigned 64-bit integer. |
| `@` | Sets the current value to the number of seconds in a UNIX-timestamp representation of the current time. |
| `@@` | Sets the current value to the number of nanoseconds in a UNIX-timestamp representation of the current time. |
| `$` | Uses the first of two stacks when calling further commands that make use of a stack. |
| `$$` | Uses the second of two stacks when calling further commands that make use of a stack. |
| `:` | Pushes the current value to the current stack. |
| `;` | Sets the current value to a value popped from the current stack. |
| `#` | Pops all the values from the current stack, performs an 8-bit hash on them and sets the current value to the result. |
| `##` | Pops all the values from the current stack, performs a 64-bit hash on them and sets the current value to the result. |
| `.` | Saves the current stack to a file, using the Unicode/ASCII representation of each value on the stack. The filename is based on the current value. |
| `,` | Loads the current stack from a file, using the Unicode/ASCII representation of each value on the stack. The filename is based on the current value. |
| `\|` | Deletes a file representing a saved stack. The filename is based on the current value. |
| `\|\|` | Clears the current stack. |
| `%\|` | Resets all state (i.e. clears both of the stacks and sets the current value to `0`). |
| ` ` | Whitespace can be used to separate two single-character commands that could otherwise be interpreted as a multi-character command. |
| `(` ... `)` | Creates a new context with a new current value of zero, in which any commands between the brackets are called, then adds the current value of the created context to the current value of the surrounding context. |
| `[` ... `]` | Creates a new context with a new current value of zero, in which any commands between the brackets are called, then subtracts the current value of the created context from the current value of the surrounding context. |
| `<` ... `>` | Runs any commands between the brackets repeatedly while the current value does not equal `0`. |
| `<<` ... `>>` | Runs any commands between the brackets repeatedly while the current value equals `0`. |
| `{` ... `}` | Ignores all characters and commands between the braces, allowing for human-readable comments. |