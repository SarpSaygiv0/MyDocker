package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Program is executing")

	command := os.Args[3]
	args := os.Args[4:len(os.Args)]

	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
