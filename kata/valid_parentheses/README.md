# valid_parentheses Notes

First create package names in go files

Rename main to match the package name

Set up initial function code in main.go

Add initial test code and fix any test names

set up module

```bash
go mod init okscoring.com/example
go mod tidy
```

write your solution code in main.go

then run the tests

```bash
go test
```

## Top Answer

```go

```

## Clever

```go

```

What did I learn in this one?

I would say this was kinda a dynamic programming proglem, also kinda data structures.
Essentially the idea is that you can have these "sub problems" that are solveable in terms of the rules. Like if you have a nested set of parens, that fits the rules nicely. So solve for the sub problems and then remove them from the problem space to make it easier to solve the other sub problems.

For the stack solution this is similar, you are just tracking the openings as you go and removing if you find a match and short circuiting otherwise. A more efficient algorithm in this case than recursion.
