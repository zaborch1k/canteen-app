package main

import (
    "log"

    "canteen-app/internal/app"
)

func main() {
    a, err := app.New();
    if err != nil {
        log.Fatal(err)
    }

    log.Println("starting server on :8080")
    if err := a.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
