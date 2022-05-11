# binary_tree_complete Notes

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

Algorithm explanation

123456

1
2 3
4 5 6

pass 1
prevChildCount = 2
node = 1

left 2 - queued
right 3 - queued
childrenCount = 2

pass 2
prevChildCount = 2
node = 2
left 4 - queued
right 5 - queued
childrenCount == 2

pass 3
prevChildCount = 2
node = 3
left 6 - queued
right nil - emptyNodeFound true
childrenCount = 1

pass 4
prevChildCount = 1
node 4
left nil
right nil
childrenCount = 0

pass 5
prevChildCount = 0
node 5
left nil
right nil
childrenCount = 0

pass 6
prevChildCount = 0
node 6
left nil
right nil
childrenCount = 0
