package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func ExecFunctions() {
	p("Contains: ", s.Contains("foo", "oo"))
	p("Count: ", s.Count("Joyston", "o"))
	p("HasPrefix: ", s.HasPrefix("Joyston", "Jo"))
	p("HasSuffix: ", s.HasSuffix("Joyston", "on"))
	p("Join: ", s.Join([]string{"a", "b", "c"}, "-"))
	p("Index: ", s.Index("Joyston", "on"))
	p("Split: ", s.Split("J-o-y-s-t-o-n", "-"))
	p("Repeat: ", s.Repeat("a", 4))
	p("Replace: ", s.Replace("foo", "o", "0", 1))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("ToUpper: ", s.ToUpper("foo"))
	p("ToLower: ", s.ToLower("FOO"))
}
