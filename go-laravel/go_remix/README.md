# Go Remix

## Learning Notes

In the go module file, using this syntax

```text
replace github.com/okeefem2/celeritas => ../celeritas
```

tells go that whenever that package is asked for, return the folder contents instead. This is really cool. Basically aliasing

running the following code will create a vendor folder with a local copy of the dependencies (npm node modules lol)

```bash
go mod vendor
```

this allows us to see the updates to the local package as we develop it right away

### Parameter Deductive typing

```go
// In both of these functions, since b is an int, and we are using an operator on them both
// go is able to deduce the type of a
func TestFunc(a, b int) int {
  return a + b
}


func TestFunc2(a, b int) int {
  return a - b
}

func TestFunc3(a, b int) int {
  return a * b
}
```
