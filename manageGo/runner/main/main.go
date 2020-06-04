package main

import (
    "log"
    "os"
    "time"
    "wheel/manageGo/runner"
)

const timeout = 2 * time.Second

func CreateTask() func(int) {
    return func(id int) {
        time.Sleep(time.Second * time.Duration(id))
        log.Printf("Processor - Task #%d.", id)
    }
}

func main() {
    log.Println("Starting work...")

    r := runner.New(timeout)

    r.Add(CreateTask(), CreateTask(), CreateTask())

    if err := r.Start(); err != nil {
        switch err {
        case runner.ErrInterrupt:
            log.Println("Terminating due to interrupt...")
            os.Exit(1)
        case runner.ErrTimeout:
            log.Println("Terminating due to interrupt...")
            os.Exit(1)
        }
    }

    log.Println("Process end...")
}
