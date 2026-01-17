package main

import (
	"log"
	"os"
)

func main() {
	// 1. Basic Logging
	log.Println("This is a standard log message")

    // 2. Customizing the Logger 
    // We can add a prefix and flags (date, time, file name) 
    log.SetPrefix("APP_LOG: ")
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

    log.Println("This log shows the file name and line number!")

    // 3. Log to a File (Instead of the terminal)
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Failed to open log file")
    }

    log.SetOutput(file)
    log.Println("This message is hidden from the terminal, it's in app.log!" )

    // 4. Fatal vs Panic (Warning: These stop the program)
    log.Panic("This prints the log + a stack trace and stops the app")
    // log.Fatal("This prints the log and exits with status 1 immediately")
}

