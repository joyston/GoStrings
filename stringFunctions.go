package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func ExecFunctions() {
	p("Contains: ", strings.Contains("foo", "oo"))
	p("ToUpper: ", strings.ToUpper("foo"))
}
