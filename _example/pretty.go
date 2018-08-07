package main

import (
	"fmt"

	"github.com/qjpcpu/go-prettyjson"
)

func main() {
	v := map[string]interface{}{
		"str":   "foo",
		"conn":  "root:root@tcp(10.0.2.2:3306)/rollcandy?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai",
		"num":   100,
		"bool":  false,
		"null":  nil,
		"array": []string{"foo", "bar", "baz"},
		"map": map[string]interface{}{
			"foo": "bar",
		},
	}
	s, _ := prettyjson.Marshal(v)
	fmt.Println(string(s))
}
