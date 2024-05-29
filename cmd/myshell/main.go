package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var built_in_cmnds = []string{"echo", "type", "exit"}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		cmnd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmnd = strings.TrimSpace(cmnd)

		if strings.Contains(cmnd, "type") {
			cmnd_t := strings.Split(cmnd, " ")
			type_chck := cmnd_t[1]
			if slices.Contains(built_in_cmnds, type_chck) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", type_chck)
			} else {
				fmt.Fprintf(os.Stdout, "%s not found\n", type_chck)
			}
			continue
		}

		if strings.Contains(cmnd, "echo") {
			var response strings.Builder
			text := strings.Split(cmnd, " ")
			for i, el := range text {
				if i == 0 {
					continue
				}
				response.WriteString(el)
				response.WriteString(" ")
			}
			fmt.Fprintf(os.Stdout, "%s\n", strings.TrimSpace(response.String()))
			continue
		}
		if strings.Contains(cmnd, "exit") {
			// exit:=strings.Split(cmnd," ")
			// exit_code:=exit[0]
			break
		}

		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmnd)
	}
}
