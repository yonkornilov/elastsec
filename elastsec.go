package main

import (
    "fmt"
    "./internal/beats/auth"
    "./internal/beats/filechange_attempt"
    "./internal/event"
)

func main() {
    eventBus := make(chan event.Event)

    go auth.Loop(eventBus)
    go filechange_attempt.Loop(eventBus)

    for event := range eventBus {
        fmt.Printf("New `%s` event on host `%s` at `%s`: %s\n",
        event.Type, event.Beat.Host, event.Time, event.Message)
    }
}
