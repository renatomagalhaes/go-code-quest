package main

import (
    "fmt"
    "os"

    "go-code-quest/levels"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "goquest",
    Short: "Go Code Quest - A CLI game to learn Go",
}

var startCmd = &cobra.Command{
    Use:   "start [level]",
    Short: "Start a new level",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        level := args[0]
        switch level {
        case "level1":
            levels.StartLevel1()
        case "level2":
            levels.StartLevel2()
        default:
            fmt.Println("Level not found")
        }
    },
}

func main() {
    rootCmd.AddCommand(startCmd)
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
