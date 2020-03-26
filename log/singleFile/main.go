package main

// 该包展示了如何进行日志记录
import (
    "log"
    "os"
)

var logger *log.Logger

// for logging
func info(args ...interface{}) {
    logger.SetPrefix("INFO ")
    logger.Println(args...)
}

func danger(args ...interface{}) {
    logger.SetPrefix("ERROR ")
    logger.Println(args...)
}

func warning(args ...interface{}) {
    logger.SetPrefix("WARNING ")
    logger.Println(args...)
}

func main() {
    file, err := os.OpenFile("dispatcher.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file", err)
    }
    logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

    info("hello world!")
    warning("hello world!")
    danger("hello world!")
}
