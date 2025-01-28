package cmd

import (
    "fmt"
    "os"
)

func Execute() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: github-activity <username>")
        os.Exit(1)
    }

    username := os.Args[1]
    err := FetchActivity(username)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
