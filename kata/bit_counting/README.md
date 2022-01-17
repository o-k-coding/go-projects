# Notes

First created library file

then created test files

the suite file sets up the test runner

Then created the module

```bash
go mod init okscoring.com/bit_counting
go mod tidy
```

then ran the tests

```bash
go test
```

Ran into a failure

./count_bits_test.go:10:11: undefined: CountBits
./count_bits_test.go:11:11: undefined: CountBits
./count_bits_test.go:12:11: undefined: CountBits
./count_bits_test.go:13:11: undefined: CountBits
./count_bits_test.go:14:11: undefined: CountBits
FAIL okscoring.com/bit_counting [build failed]

had to make sure to import the package into the test file

```go
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "okscoring.com/bit_counting"
)
```

## Top Answer

The math package seems to have a function for counting the number of ones in the bits of a number

```go
package kata

import "math/bits"

func CountBits(n uint) int {
  return bits.OnesCount(n)
}
```

## Clever

A clever answer that assigns the bits function to a variable CountBits which the tests would call

```go
package kata
import "math/bits"
var CountBits = bits.OnesCount
```
