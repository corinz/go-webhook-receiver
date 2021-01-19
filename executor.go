package webhook

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Executor struct executes stores arrays of execution info
//  Executables []string | Usage: single word executable statements or executable files only
//  LogicTests []string  | Usage: myParameter "{eq,ne,lt,gt}" <string> OR <int>
//  TestEnabled []int    | Usage: Corresponds to active/inactive executables. 0=disabled, 1=enabled */
type Executor struct {
	Executables []string
	LogicTests  []string
	TestEnabled []int
}

// Add Executable, file or 0 arg command
func (e *Executor) Add(cmd string, logic string) {
	e.Executables = append(e.Executables, cmd)
}

// Execute array or commands/files preprended with /bin/bash -c
func (e *Executor) Execute() error {

	//for loop and os call on Executables
	for i, v := range e.Executables {
		if e.TestEnabled[i] == 0 {
			continue // skip disabled tests
		}

		cmdArr := strings.Split(v, " ")

		// /bin/bash -c ...
		cmd := exec.Command("./exec.sh", cmdArr...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println("Error: ", err)
			return err
		}
	}
	return nil
}

// Print executable files, or commands
func (e *Executor) Print() {
	for i, v := range e.Executables {
		fmt.Println("Arg ", i, v)
	}
}
