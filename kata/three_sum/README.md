# three_sum Notes

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

Needed some help with this one <https://www.code-recipe.com/post/three-sum>

Notes

Sort the array up front, this will group duplicate elements and allow us to ignore them without use of a hashmap.
Then use two loops to process the triplets.

Fix the outer loop, and use two pointers (indexes) in the inner loop.
The two pointers in the inner loop will move inwards towards each other.

so example -1, -1, 0, 1, 2

each column shows a loop, the left column is the first (outer) loop
-1, -1, 2 - sum == 0 so we store triplet and decrement the 3rd index
-1, -1, 1 - sum < 0 so we increment the 2nd index
-1, 0, 1 - sum == 0 so we store triplet and decrement 3rd until we get to a number that isn't the same as the one we just saw in the 3rd spot
-1, 0, 0 - 3rd is not > 2nd anymore, so break inner loop and increment 1st
-1 - before starting the inner loop, we see that this value is the same as the previous 1st index value, so we skip and increment
0, 1, 2 - sum is > 0 so we decrement 3rd
0, 1, 1 - 2nd and 3rd are equal so we stop

No more numbers for the 1st loop to process (loops until length - 3) so we are done!


## Top Answer

```go

```

## Clever

```go

```
