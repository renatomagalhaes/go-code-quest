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
        case "level3":
            levels.StartLevel3()
        case "level4":
            levels.StartLevel4()
        case "level5":
            levels.StartLevel5()
        case "level6":
            levels.StartLevel6()
        case "level7":
            levels.StartLevel7()
        case "level8":
            levels.StartLevel8()
        case "level9":
            levels.StartLevel9()
        case "level10":
            levels.StartLevel10()
        case "level11":
            levels.StartLevel11()
        case "level12":
            levels.StartLevel12()
        case "level13":
            levels.StartLevel13()
        case "level14":
            levels.StartLevel14()
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
