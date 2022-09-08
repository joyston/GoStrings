package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func ExecFormatting() {
	p := point{1, 2}

	fmt.Printf("struc1: %v\n", p)
	fmt.Printf("struc2: %+v\n", p)
	fmt.Printf("struc2: %#v\n", p)
	fmt.Printf("type: %T\n", p)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 123)
	fmt.Printf("binary: %b\n", 12)
	fmt.Printf("character: %c\n", 33)
	fmt.Printf("hex: %x\n", 145)
	fmt.Printf("float1: %f\n", 12.3)
	fmt.Printf("scientific float: %e\n", 12000000.3)
	fmt.Printf("scientific float: %e\n", 123400000.0)
	fmt.Printf("str1: %s\n", "\"string\"")
	fmt.Printf("str2: %q\n", "\"string\"")
	fmt.Printf("str3: %x\n", "hex this")
	fmt.Printf("Pointer: %p\n", &p)
	fmt.Printf("Width1: |%6d|%6d|\n", 12, 124)
	fmt.Printf("Width2: |%6.2f|%6.2f|\n", 1.2, 1.243)
	fmt.Printf("Width3: |%-6.2f|%-6.2f|\n", 1.2, 1.243)
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
	fmt.Printf("width4: |%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
