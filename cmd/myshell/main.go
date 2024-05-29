package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

var built_in_cmnds = []string{"echo", "type", "exit"}

func checkInPath(name string) (string, bool) {
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, el := range paths {
		fullpath := filepath.Join(el, name)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath, true
		}
	}
	return "", false
}

func main() {

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
			} else if path, ok := checkInPath(type_chck); ok {
				fmt.Fprintf(os.Stdout, "%s is %s\n", type_chck, path)
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
		} else {
			exe := strings.Split(cmnd, " ")

			if _, err := exec.LookPath(exe[0]); err == nil {
				// use(path)
				cmd := exec.Command(exe[0], exe[1])
				if out, err := cmd.CombinedOutput(); err == nil {
					fmt.Fprintf(os.Stdout, "%s", out)
				} else {
					log.Fatal(err)
				}
			} else {
				fmt.Fprintf(os.Stdout, "%s: command not found\n", cmnd)
			}
			continue
		}
	}
}
