package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello world")
	fmt.Fprintln(os.Stdout, "welcome")
	io.WriteString(os.Stdout, "Hiiiiiiii")
}
