package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	arg := os.Args
	var arg_open string
	if len(arg) == 2 {
		arg_open = arg[1]
	}

	app := "/usr/bin/flatpak"
	args := []string{"run", "--branch=stable", "--arch=x86_64", "--command=code", "--file-forwarding", "com.visualstudio.code", "--reuse-window", arg_open}

	program := exec.Command(app, args...)

	err := program.Start()

	if err != nil {
		log.Printf("Error to run program %s: %s\n", err, strings.Join(program.Args, " "))
	}

	fmt.Printf("Executing %s %s with %d PID\n", app, args[5], program.Process.Pid)

}
