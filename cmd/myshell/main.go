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
  
  if strings.Contains(cmnd,"exit"){
    // exit:=strings.Split(cmnd," ")
    // exit_code:=exit[0]
    break
  }
  var response strings.Builder
  if strings.Contains(cmnd,"echo"){
    text:=strings.Split(cmnd," ")
    for i,el:= range text{
      if i==0 {continue}
      response.WriteString(el)
      response.WriteString(" ")
    }
    fmt.Fprintf(os.Stdout,"%s\n",strings.TrimSpace(response.String()))
  }else{
      fmt.Fprintf(os.Stdout,"%s: command not found\n",cmnd)
    }


  }
}
