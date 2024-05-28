package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
  for ; ;{
  fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
  cmnd,_:=bufio.NewReader(os.Stdin).ReadString('\n')
  cmnd = strings.TrimSpace(cmnd)
  fmt.Fprintf(os.Stdout,"%s: command not found\n",cmnd)
  }
}
