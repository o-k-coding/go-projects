# Notes

First create package names in go files

Set up initial function code in main.go

Add initial test code and fix any test names

set up module

```bash
go mod init okscoring.com/bit_counting
go mod tidy
```

write your solution code in main.go

then run the tests

```bash
go test
```

## Top Answer

Using a string replacer. This is cool and good to know about.

```go
package kata

import "strings"

var dnaReplacer *strings.Replacer = strings.NewReplacer(
  "A", "T",
  "T", "A",
  "C", "G",
  "G", "C",
)

func DNAStrand(dna string) string {
  return dnaReplacer.Replace(dna)
}
```

## Clever

```go

```
