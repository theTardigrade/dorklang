# dorklang

This is an interpreter, written in Go, for **dorklang**, which is an esoteric programming language devised by me.

## Examples

After cloning the repository to your local machine, try running one of the example programs with the commands below:

```
cd interpreter
go run ./flag.go ./main.go --file=../examples/countdown.dork
```

## Storage

### Current Value

Each **dorklang** program has access to a 64-bit unsigned integer known as the **current value**, which is automatically assigned the value `0` when the program begins.

It is possible to enter a new context and gain access to another current value.

Given that integer values rollover, it is possible to reach the maximum value that can be held by the current value by setting it to `0`, if it isn't already, and then subtracting `1`.

### Current Stack

There are also two stacks available for storage. Only one of these is set as the **current stack** at any one time.

The current value can be pushed onto and popped from the current stack. Each stack can hold a maximum of `1_048_576` values, each of which is a 64-bit unsigned integer.

Only one pair of stacks is available throughout the lifetime of the program, even if a new context is entered.

## Syntax

Below is an overview of all the commands that can be used in **dorklang** source-code files:

| Command | Function |
| :--------: | ------- |
| `+` | Adds `1` to the **current value**. |
| `++` | Adds `8` to the **current value**. |
| `%+` | Pops the two topmost values from the **current stack**, adds one to the other and sets the **current value** to the result. | 
| `%++` | Pops all of the values from the **current stack**, adds each of them to the others and sets the **current value** to the result. | 
| `-` | Subtracts `1` from the **current value**. |
| `--` | Subtracts `8` from the **current value**. |
| `%-` | Pops the two topmost values from the **current stack**, subtracts one from the other and sets the **current value** to the result. | 
| `%--` | Pops all of the values from the **current stack**, subtracts each of them from the others and sets the **current value** to the result. | 
| `/` | Divides the **current value** by `2`. |
| `//` | Divides the **current value** by `8`. |
| `%/` | Pops the two topmost values from the **current stack**, divides one from the other and sets the **current value** to the result. | 
| `%//` | Pops all of the values from the **current stack**, divides each of them from the others and sets the **current value** to the result. | 
| `*` | Multiplies the **current value** by `2`. |
| `**` | Multiplies the **current value** by `8`. |
| `%*` | Pops the two topmost values from the **current stack**, multiplies one with the other and sets the **current value** to the result. | 
| `%**` | Pops all of the values from the **current stack**, multiplies each of them with the others and sets the **current value** to the result. | 
| `^` | Squares the **current value** (i.e. multiplies it by itself). |
| `^^` | Cubes the **current value** (i.e. multiplies it by itself twice). |
| `!` | Prints the **current value** to the screen as a Unicode/ASCII character. |
| `!!` | Prints the **current value** to the screen as a decimal number. |
| `?` | Waits for a Unicode/ASCII character to be given as input, then sets the **current value** to its numerical value. |
| `??` | Waits for a decimal number to be given as input, then sets the **current value** to it. |
| `~` | Sets the **current value** to `0`. |
| `'` | Sets the **current value** to the size of a byte (i.e. `8`). |
| `''` | Sets the **current value** to the size of eight bytes (i.e. `64`). |
| `"` | Sets the **current value** to the size of a kibibyte (i.e. `8_192`). |
| `""` | Sets the **current value** to the size of eight kibibytes (i.e. `65_536`). |
| `%'` | Sets the **current value** to the size of a mebibyte (i.e. `8_388_608`). |
| `%''` | Sets the **current value** to the size of eight mebibytes (i.e. `67_108_864`). |
| `%"` | Sets the **current value** to the size of a gibibyte (i.e. `8_589_934_592`). |
| `%""` | Sets the **current value** to the size of eight gibibytes (i.e. `68_719_476_736`). |
| `` ` `` | Sets the **current value** to a random number between `0` and `255`. |
| ``` `` ``` | Sets the **current value** to a random number between `0` and the maximum value for an unsigned 64-bit integer. |
| `@` | Sets the **current value** to the number of seconds in a UNIX-timestamp representation of the current time. |
| `@@` | Sets the **current value** to the number of nanoseconds in a UNIX-timestamp representation of the current time. |
| `%&` | Performs a logical AND operation on the two topmost values in the **current stack**, setting the **current value** to `1` if both values from the stack do not equal `0`, otherwise setting the **current value** to `0`. No values are popped from the **current stack**. |
| `%&&` | Performs a logical AND operation on all of the values in the **current stack**, setting the **current value** to `1` if all values from the stack do not equal `0`, otherwise setting the **current value** to `0`. No values are popped from the **current stack**. |
| `\` | Inverts the **current value** as though it were a boolean (i.e. sets the **current value** to `0` if it is not already `0`, otherwise sets it to `1`). |
| `$` | Uses the first of two stacks when calling further commands that make use of a stack. |
| `$$` | Uses the second of two stacks when calling further commands that make use of a stack. |
| `:` | Pushes the **current value** to the end of the **current stack**. |
| `%:` | Sets the **current value** to the number of values stored in the **current stack**. |
| `;` | Sets the **current value** to a value popped from the end of the **current stack**. |
| `%;` | Sets the **current value** to a value popped from a random position in the **current stack**. |
| `#` | Pops all the values from the **current stack**, performs an 8-bit hash on them and sets the **current value** to the result. |
| `##` | Pops all the values from the **current stack**, performs a 64-bit hash on them and sets the **current value** to the result. |
| `s` | Sorts the **current stack** in ascending order, so that the largest values are at the top and the smallest values are at the bottom. |
| `ss` | Sorts the **current stack** in descending order, so that the largest values are at the bottom and the smallest values are at the top.
| `%s` | Shuffles the **current stack** so that the values are in a random order. |
| `x` | Swaps the top two values on the **current stack**, so that the topmost becomes the second-to-topmost (and *vice versa*). |
| `r` | Reverses the order of all values in the **current stack**. |
| `i` | Pushes an iota-range of values to the **current stack**, from `0` inclusive to the **current value** exclusive. |
| `ii` | Pushes an iota-range of values to the **current stack**, from `1` inclusive to the **current value** exclusive. |
| `.` | Saves the **current stack** to a file, using the Unicode/ASCII representation of each value on the stack. The filename is based on the **current value**. |
| `,` | Loads the **current stack** from a file, using the Unicode/ASCII representation of each value on the stack. The filename is based on the **current value**. |
| `\|` | Deletes a file representing a saved stack. The filename is based on the **current value**. |
| `\|\|` | Clears the **current stack**. |
| `%\|` | Resets all state (i.e. clears both of the stacks and sets the **current value** to `0`). |
| ` ` | Whitespace can be used to separate two single-character commands that could otherwise be interpreted as a multi-character command. |
| `(` ... `)` | Creates a new context with a new **current value** of zero, in which any commands between the brackets are called, then adds the **current value** of the created context to the **current value** of the surrounding context. |
| `((` ... `))` | Creates a new context with a new **current value** of zero, in which any commands between the brackets are called, then multiplies the **current value** of the created context by the **current value** of the surrounding context. |
| `[` ... `]` | Creates a new context with a new **current value** of zero, in which any commands between the brackets are called, then subtracts the **current value** of the created context from the **current value** of the surrounding context. |
| `[[` ... `]]` | Creates a new context with a new **current value** of zero, in which any commands between the brackets are called, then divides the **current value** of the surrounding context by the **current value** of the created context. |
| `<` ... `>` | Runs any commands between the brackets repeatedly while the **current value** does not equal `0`. |
| `<<` ... `>>` | Runs any commands between the brackets repeatedly while the **current value** equals `0`. |
| `{` ... `}` | Ignores all characters and commands between the braces, allowing for human-readable comments. |
| `{{` ... `}}` | Reads one or more files. The names of the files are given between the braces, separated by whitespace. If a file has a `.dork` extension, the commands it contains are run by the interpreter, but if a file has any other extension, the contents of the file is pushed onto the **current stack**. All commands within the braces are ignored. |

## Support

Please consider donating if you want to support my work:

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/S6S2EIRL0)