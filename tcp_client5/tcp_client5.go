package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	// Authentication
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	fmt.Fprint(conn, username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	fmt.Fprint(conn, password)

	// Display server's response (authentication result)
	authResponse, _ := serverReader.ReadString('\n')
	fmt.Print(authResponse)

	// Start a goroutine to listen for messages from the server
	go func() {
		for {
			message, err := serverReader.ReadString('\n')
			if err != nil {
				fmt.Println("Server disconnected.")
				return
			}
			fmt.Print(message)
		}
	}()

	// Loop to send messages to the server
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			fmt.Println("Exiting...")
			break
		}

		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}
