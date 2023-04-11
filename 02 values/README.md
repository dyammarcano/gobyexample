# Go by Example: Values

````go
// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.

package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang")

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
````

Strings, which can be added together with +.

Integers and floats.

Booleans, with boolean operators as you’d expect.

```bash
$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```

### Run example

[Example](https://goplay.tools/snippet/fF9Rv-9DHJu)
