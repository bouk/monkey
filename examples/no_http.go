package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/bouk/monkey"
)

func main() {
	monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient), "send", func(_ *http.Client, _ *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("no requests allowed")
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err)
}
