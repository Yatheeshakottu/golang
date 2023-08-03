package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    fmt.Println("Enter a paragraph:")
    reader := bufio.NewReader(os.Stdin)
    paragraph, _ := reader.ReadString('\n')
    fmt.Printf("The paragraph contains %d characters.\n", len(paragraph)-1)
}
/*In this program, we use the bufio package to read input from the user and the os package to read input from standard input. We then use the len() function to count the number of characters in the paragraph, subtracting 1 to account for the newline character at the end of the input. Finally, we print out the number of characters using fmt.Printf().*/a



