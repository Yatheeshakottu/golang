package main

import (
	"fmt"
)

func main() {
	jb := []string{"james Bond,chacolate,vennila"}
	fmt.Println(jb)
	mb := []string{"miss monneypenny,strawberry"}
	fmt.Println(mb)
	xp := [][]string{jb, mb}
	fmt.Println(xp)
}
