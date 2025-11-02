package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: tiny-docker run <command> [args...]")
	os.Exit(1)
    }

    switch os.Args[1] {
    case "run":
        run()
    default:
        fmt.Printf("Unknown command: %s\n", os.Args[1])
	os.Exit(1)
    }
}

func run() {
    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Error running command:", err)
    }
}

