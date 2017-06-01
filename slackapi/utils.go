package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// PrintAndExit prints the JSON-encoded data as a string and exits.
func PrintAndExit(data interface{}) {
	PrintFormattedJSON(data)
	os.Exit(0)
}

// PrintFormattedJSON prints the JSON-encoded data as a string.
func PrintFormattedJSON(data interface{}) {
	response, err := json.MarshalIndent(data, "", "\x20\x20")

	if err != nil {
		fmt.Println("error;", err)
	}

	fmt.Printf("%s\n", response)
}

// PrintInlineJSON prints the JSON-encoded data as a formatted string.
func PrintInlineJSON(data interface{}) {
	response, err := json.Marshal(data)

	if err != nil {
		fmt.Println("error;", err)
	}

	fmt.Printf("%s\n", response)
}

// ShellExec executes an external command and returns its output.
func ShellExec(kommand string) []byte {
	var binary string
	var parts []string
	var arguments []string

	if kommand == "" {
		fmt.Println("error; invalid empty command")
		return []byte{}
	}

	parts = strings.Fields(kommand)
	binary = parts[0]
	arguments = parts[1:]

	response, err := exec.Command(binary, arguments...).Output()

	if err != nil {
		fmt.Println("error;", err)
	}

	return bytes.Trim(response, "\n")
}
