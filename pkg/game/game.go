package game

import "fmt"

type Level interface {
    Start()
}

func StartLevel(level Level) {
    fmt.Println("Starting level...")
    level.Start()
}
