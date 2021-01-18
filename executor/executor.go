package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Executor struct
type Executor struct {
	Executables []string
}

// Add Executable, file or 0 arg command
func (e *Executor) AddExecutable(cmd string) {
	e.Executables = append(e.Executables, cmd)
}

// Execute array or commands/files preprended with /bin/bash -c
func (e *Executor) Execute() error {
	//for loop and os call on Executables
	for _, v := range e.Executables {
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
