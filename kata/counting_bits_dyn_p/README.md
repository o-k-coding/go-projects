# counting_bits_dyn_p Notes

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

func countBits(n int) []int {
    arr := make([]int, n+1)
    for i:=0;i<=n;i++ {
        arr[i]=countSetBits(i)
}
return arr
}

func countSetBits(n int) int {
    count:=0
    for n>0 {
        count+=n&1
        n=n>>1
}
return count
}

```

## Clever

```go

```
