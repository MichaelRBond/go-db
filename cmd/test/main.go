package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Running tests...")

	cmd := exec.Command("go", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Tests failed:", err)
		os.Exit(1)
	}

	fmt.Println("All tests passed!")
}
