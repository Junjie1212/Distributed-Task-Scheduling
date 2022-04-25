package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
	)

	// cmd = exec.Command("/bin/bash", "-c", "echo 1;echo2;")

	cmd = exec.Command("C:\\cygwin64\\bin\\bash.exe", "-c", "echo 1")

	// The returned error is nil if the command runs
	err = cmd.Run()

	fmt.Println(err)
}
