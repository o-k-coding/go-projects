# Notes

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

Just uses a for loop. Certainly more concise than mine.

```go
package kata

func Tribonacci(signature [3]float64, n int) (r []float64) {
  r = signature[:]
  for i := 0; i < n; i++ {
    l := len(r)
    r = append(r, r[l-1]+r[l-2]+r[l-3])
  }
  return r[:n]
}
```

## Clever

```go

```
