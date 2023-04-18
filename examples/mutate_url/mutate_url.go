package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://www.bing.com/search?q=java&safe=active"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()

	fmt.Printf("mutate query: %s", u.String())
}
