# cleaner-go

# cleaner.go
Writing cleaner Go code with examples

1. Comments.
    1.1 Don't just write some comment

2. Function return arguments.
    2.1. Don't name return arguments on functions.
    2.2. Don't return an error as an argument that isn't the final argument.
    2.3. If returning an error and another value, dont return a non-nil value if there is an error.

3. Variable names.
    3.1 Don't use hungarian names which include the variables type.

4. Conditional statements.
	4.1 Dont put long conditional statements all on one line.
    4.2 Don't use multiple if statements back to back if only one is required.

5. Function Declaration
    5.1 Dont declare a function inside a function call.

6. Structs
    6.1 Don't change values on structs immediately after constructing them.

7. Logic
    7.1 Don't use fuzzy logic.


# cleaner_test.go
Writing cleaner Go test code with examples

1. Table tests.
2. Don't test the standard library.