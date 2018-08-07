package main

import (
	"fmt"

	"github.com/qjpcpu/go-prettyjson"
)

func main() {
	v := map[string]interface{}{
		"str":   "foo",
		"num":   100,
		"bool":  false,
		"null":  nil,
		"array": []string{"foo", "bar", "baz"},
		"map": map[string]interface{}{
			"foo": "bar",
		},
	}
	f := prettyjson.NewFormatter()
	f.DisabledColor = true
	s, _ := f.Marshal(v)
	fmt.Println(string(s))
}
