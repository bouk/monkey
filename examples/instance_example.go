package main

import (
	"fmt"
	"net"
	"net/http"
	"reflect"

	"github.com/bouk/monkey"
)

func main() {
	var d *net.Dialer
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Dial", func(_ *net.Dialer, _, _ string) (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
}

// This will do same as done in main:
func PatchInstanceMethodFlexibleExample1() {
	var d *net.Dialer
	monkey.PatchInstanceMethodFlexible(reflect.TypeOf(d), "Dial", func() (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
}

// Eliminate/use input variables example:
func PatchInstanceMethodFlexibleExample2() {
	var d *net.Dialer
	monkey.PatchInstanceMethodFlexible(reflect.TypeOf(d), "Dial", func(_, _ interface{}, address string) (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed for address: %s", address)
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
}