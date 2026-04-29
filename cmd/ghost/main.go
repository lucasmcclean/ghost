package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/lucasmcclean/ghost/internal/runtime"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: ghost <container|pid>")
		os.Exit(1)
	}

	pid, err := runtime.Resolve(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("PID:", pid)

	cmd := exec.Command("nsenter",
		"-t", strconv.Itoa(pid),
		"-m", "-u", "-i", "-n",
		"--", "/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
