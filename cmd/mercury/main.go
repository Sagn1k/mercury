package main

import (
    "log"
    "net"
    "bufio"
    "strings"
)

func main() {
    smtpAddr := "127.0.0.1:2525"
    log.Println("Starting SMTP server at", smtpAddr)
    listener, err := net.Listen("tcp", smtpAddr)
    if err != nil {
        log.Fatalf("Failed to start SMTP server: %v", err)
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        line := scanner.Text()
        log.Println("Received:", line)

        // Handle various SMTP commands
        // Respond to the client based on the SMTP protocol
        if strings.HasPrefix(line, "HELO") {
            conn.Write([]byte("250 Hello\r\n"))
        } else if strings.HasPrefix(line, "MAIL FROM:") {
            conn.Write([]byte("250 Ok\r\n"))
        } else if strings.HasPrefix(line, "RCPT TO:") {
            conn.Write([]byte("250 Ok\r\n"))
        } else if strings.HasPrefix(line, "DATA") {
            conn.Write([]byte("354 End data with <CR><LF>.<CR><LF>\r\n"))
        } else if strings.HasPrefix(line, ".") {
            conn.Write([]byte("250 Ok: queued as 12345\r\n"))
        } else if strings.HasPrefix(line, "QUIT") {
            conn.Write([]byte("221 Bye\r\n"))
            break
        } else {
            conn.Write([]byte("500 Command not recognized\r\n"))
        }
    }

    if err := scanner.Err(); err != nil {
        log.Printf("Error reading from client: %v", err)
    }
}