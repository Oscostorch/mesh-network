package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

// In-memory "database" of users and passwords
var validCredentials = map[string]string{
	"user1": "pass1",
	"user2": "pass2",
	"user3": "pass3",
	"user4": "pass4",
	"user5": "pass5",
}

var (
	clients     = make(map[net.Conn]string) // Map to track connected clients and their usernames
	activeUsers = make(map[string]bool)     // Set to track active usernames
	messages    []string                    // Slice to store broadcasted messages
	mu          sync.Mutex                  // Mutex for concurrent access to maps and messages slice
)

func main() {
	startServer()
}

func startServer() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server is listening on port 8081...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		mu.Lock()
		username := clients[conn]
		delete(clients, conn)         // Remove client on disconnect
		delete(activeUsers, username) // Remove username from active users
		mu.Unlock()
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	authenticated := false
	var username string

	// Authentication loop
	for !authenticated {
		fmt.Fprint(conn, "Enter username: ")
		usernameInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(conn, "Error reading username:", err)
			return
		}
		username = strings.TrimSpace(usernameInput)

		fmt.Fprint(conn, "Enter password: ")
		passwordInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(conn, "Error reading password:", err)
			return
		}
		password := strings.TrimSpace(passwordInput)

		mu.Lock()
		if validPassword, exists := validCredentials[username]; exists && validPassword == password {
			if _, isActive := activeUsers[username]; isActive {
				fmt.Fprintln(conn, "This username is already logged in. Please use a different username.")
			} else {
				activeUsers[username] = true // Mark username as active
				clients[conn] = username     // Store the username for the client
				fmt.Fprintf(conn, "Welcome, %s!\n", username)
				authenticated = true

				// Send message history to the new client
				fmt.Fprintln(conn, "Message history:")
				for _, msg := range messages {
					fmt.Fprintln(conn, msg)
				}
			}
		} else {
			fmt.Fprintln(conn, "Invalid credentials. Please try again.")
		}
		mu.Unlock()
	}

	// Broadcast loop
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected.\n", username)
			break
		}

		message = strings.TrimSpace(message)
		fmt.Printf("[%s]: %s\n", username, message)

		broadcastMessage(conn, fmt.Sprintf("[%s]: %s", username, message))
	}
}

func broadcastMessage(sender net.Conn, message string) {
	mu.Lock()
	defer mu.Unlock()

	// Store the message in the history
	messages = append(messages, message)

	// Broadcast the message to all clients except the sender
	for client := range clients {
		if client != sender {
			_, err := client.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
		}
	}
}
