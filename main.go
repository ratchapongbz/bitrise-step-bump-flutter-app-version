package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("This is the value specified for the input 'example_step_input':", os.Getenv("example_step_input"))

	fmt.Println("This is the value specified for the input 'pubspec_path':", os.Getenv("pubspec_path"))

	fmt.Println("This is the value specified for the input 'bumping_type':", os.Getenv("bumping_type"))

	//
	// --- Step Outputs: Export Environment Variables for other Steps:
	// You can export Environment Variables for other Steps with
	//  envman, which is automatically installed by `bitrise setup`.
	// A very simple example:
	cmdLog, cmdLogErr := exec.Command("bitrise", "envman", "add", "--key", "EXAMPLE_STEP_OUTPUT", "--value", "the value you want to share").CombinedOutput()

	if cmdLogErr != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", cmdLogErr, cmdLog)
		os.Exit(1)
	}
	// You can find more usage examples on envman's GitHub page
	//  at: https://github.com/bitrise-io/envman

	appVersion, appVersionErr := exec.Command("bitrise", "envman", "add", "--key", "PUBSPEC_APP_VERSION", "--value", "1.0.0").CombinedOutput()

	if appVersionErr != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", appVersionErr, appVersion)
		os.Exit(1)
	}

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}
