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

net package has a built in ip parser lol!

```go
package kata

import "net"
func Is_valid_ip(ip string) bool {
  res := net.ParseIP(ip)
  if res == nil {
    return false
  }
  return true
}
```

## Clever

```go

```
