package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	// grab arguments
	args := strings.Split(input, " ")

	// There is no "real" cd command it is a built in functionality of the shell
	// so we have to build it here

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// prep the command to be executed
	cmd := exec.Command(args[0], args[1:]...)

	// Set output devices
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {
	// read in user input
	reader := bufio.NewReader(os.Stdin)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not get machine's host name")
		os.Exit(1)
	}

	currUser, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not get current user")
		os.Exit(1)
	}

	for {

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not get working directory")
			os.Exit(1)
		}

		// bash like prompt
		fmt.Printf("%v@%v:%v> \n", currUser.Username, hostname, cwd)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err := execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
