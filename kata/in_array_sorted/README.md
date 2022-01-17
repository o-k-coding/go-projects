# Notes

First created library file

then created test files

the suite file sets up the test runner

Then created the module

```bash
go mod init okscoring.com/in_array_sorted
go mod tidy
```

then ran the tests

```bash
go test
```

## Top Answer

Actually seems like my solution was pretty close to others solutions. Mine was slightly easier to read I think

```go
package kata

import (
  "sort"
  "strings"
)

func InArray(array1 []string, array2 []string) []string {
  seen := make(map[string]struct{})
  result := []string{}
  for _, s1 := range array1 {
    if _, ok := seen[s1]; ok {
      continue
    }
    seen[s1] = struct{}{}
    for _, s2 := range array2 {
      if strings.Contains(s2, s1) {
        result = append(result, s1)
        break
      }
    }
  }
  sort.Strings(result)
  return result
}
```
