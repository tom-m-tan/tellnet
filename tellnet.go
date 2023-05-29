package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s host port\n", os.Args[0])
        os.Exit(1)
    }

    host := os.Args[1]
    port := os.Args[2]

    addr := fmt.Sprintf("%s:%s", host, port)
    timeout := 2 * time.Second

    conn, err := net.DialTimeout("tcp", addr, timeout)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

    conn.Close()
    fmt.Printf("Port %s is open\n", port)
}
