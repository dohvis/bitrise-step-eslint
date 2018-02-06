package main

import (
	"fmt"
	"os"
	"os/exec"
)

func getESLintPath() (bool, string) {
	isExist := true
	path, err := exec.LookPath("eslint")
	if err != nil {
		path, err = exec.LookPath("node_modules/eslint/bin/eslint.js")
		if err != nil {
			fmt.Println("First of all, install eslint")
			isExist = false
		}
	}
	return isExist, path
}

func runESLint(eslintPath string, srcDirectory string) bool {
	nodePath, err := exec.LookPath("node")
	output, err := exec.Command(nodePath, eslintPath, srcDirectory).CombinedOutput()
	fmt.Printf("Output: %s\n", string(output[:]))
	hasError := false
	if err != nil {
		hasError = true
		fmt.Printf("Error: %s\n", err)
	}
	return hasError
}

func main() {
	// fmt.Println("This is the value specified for the input 'example_step_input':", )

	srcDirectory := os.Getenv("target")

	isExist, eslintPath := getESLintPath()
	hasError := false
	exitCode := 1
	if isExist {
		hasError = runESLint(eslintPath, srcDirectory)
		if !hasError {
			exitCode = 0
		}
	}
	os.Exit(exitCode)
}
