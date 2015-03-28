package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/bouk/monkey"
)

func main() {
	var guard *monkey.PatchGuard
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient), "send", func(c *http.Client, req *http.Request) (*http.Response, error) {
		guard.Unpatch()
		defer guard.Restore()

		if req.URL.Scheme == "http" {
			return nil, fmt.Errorf("no http requests allowed")
		}

		return c.Do(req)
	})

	_, err := http.Get("http://google.com")
	fmt.Println(err)
	resp, err := http.Get("https://google.com")
	fmt.Println(resp.Status, err)
}
