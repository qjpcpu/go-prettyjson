package main

import (
	"fmt"

	"github.com/qjpcpu/go-prettyjson"
)

func main() {
	s, _ := prettyjson.Format([]byte(`{"foo":"bar"}`))
	fmt.Println(string(s))
}
