# Go by Example: 

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer`s will _not_ be run when using `os.Exit`, so
	// this `fmt.Println` will never be called.
	defer fmt.Println("!")

	// Exit with status 3.
	os.Exit(3)
}
```

```bash
$ go run exit.go
exit status 3
```

### Run example

[Example](https://goplay.tools/snippet/BCJFZgwtCAJ)