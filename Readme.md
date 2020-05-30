# GoLang

## Intro

- Designed as a system language.
- Can now be used as a general purpose language.
- Automatic garbage collection, memory clearing etc.
- Only ~25 keywords.
- Cross platform.
- Compiled language.
- Strongly typed.
- Supports type inference (can infer the variable type from declaration).
- Case sensitive

## Variables and Constants

### Declaring and Initializing variables

- Variables are statically types, a string remains a string for its lifetime
- var is used for declaration at package level, they can also be grouped together
- Go always initializes variable with zero value
- reflect can be used to check types
- Go can infer types from declaration
- Type conversion is pretty much like python
- Variables declared at package level are Global variables
- := works only inside functions and only if we're initializing the variables at the same time we declare them.
- It is necessary to use all the variables, it is not possible to declare variables without using it inside functions. If we declare a variable inside a function it is necessary to use it.
- := is used to declare new variables whereas = is used to reassign values or simply assign values to a variable
- Constants are immutable, declared with const, cannot use := with constants.
- Environment variables are accessible with os package

## Pointers

- Go passes arguments by value, and not by reference.
- When we create a variable, Go creates a space for it in the memory, when we pass this variable into a function, it makes a copy of this variable with same value in different memory space, the changes are thus made only to the copy of the variable.
- & gives the memory address of a variable (the pointer value)
- & references a pointer and * de-references a pointer
- We can use pointers to pass value by reference



## Conditionals

- If accepts some statements also before the boolean expression
- Any variables declared with if exist in only that if and else if blocks and are garbage collected after that
- 