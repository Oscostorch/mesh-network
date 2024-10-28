Mesh Network Bootstrapping in Go

a basic program in Go that establishes a mesh network between two devices located on different networks.



Peer-to-Peer Communication Application in Go

This Go application facilitates secure, direct communication between two devices (Device A and Device B) over the internet. The system implements a discovery mechanism using a signaling server and a direct peer-to-peer communication setup for message exchange. It also includes basic data transmission and error handling capabilities.
Table of Contents

    Architecture
    Dependencies
    Installation
    Running the Application
    Testing the Application
    Additional Notes

Architecture

The application is divided into the following key components:

    Signaling Server:
        Role: Facilitates the discovery process by registering devices and providing a list of connected devices.
        Endpoint:
            /register: Allows a device to register itself with a unique ID and IP address.
            /devices: Provides a list of registered devices for discovery.

    Device Code:
        Role: Each device registers with the signaling server, retrieves the list of other devices, and initiates peer-to-peer communication.
        Communication: Devices use the server’s device list to establish TCP connections with each other.

    TCP Server:
        Role: Provides a centralized TCP server where clients connect, authenticate, and communicate.
        Authentication: Implements a username-password mechanism for added security.
        Message Broadcasting: Sends messages to all connected clients except the sender.

    TCP Clients:
        Role: Connects to the TCP server, authenticates with a username and password, and facilitates message exchange.
        Communication: Listens for messages from other clients and displays them on the console.

Dependencies

Ensure you have Go installed (version 1.15 or higher). You can install it from Go's official website.
Installation

    Clone the repository:

    bash

git clone https://github.com/your-repo-name
cd your-repo-name

Navigate to each module directory and build the binaries:

bash

    cd signaling_server
    go build -o signaling_server
    cd ../device
    go build -o device
    cd ../tcp_server
    go build -o tcp_server
    cd ../tcp_clients
    go build -o tcp_client1

Running the Application

To set up the system, follow these steps:
1. Start the Signaling Server

bash

./signaling_server

    The signaling server will listen on localhost:8080 for device registration and listing requests.

2. Register Devices

On each device, start the device executable:

bash

./device

    This will register the device with the signaling server and display a list of connected devices.

3. Start the TCP Server

In a separate terminal, start the TCP server:

bash

./tcp_server

    The TCP server will listen on localhost:8081 for incoming client connections.

4. Start the TCP Clients

Run multiple instances of the TCP client (e.g., tcp_client1, tcp_client2), each representing a unique client.

bash

./tcp_client1

    Each client will prompt for a username and password to authenticate with the server.

5. Sending and Receiving Messages

    Each client can send messages to the TCP server, which will broadcast them to other connected clients. Type your message and press Enter to send.

6. Termination

    To terminate a client, type exit.

Testing the Application
Discovery Mechanism

    Run multiple devices using the device executable to verify each can register with the signaling server and retrieve a list of connected devices.

Peer-to-Peer Communication

    Ensure each client connected to the TCP server can send and receive messages from other clients.

Basic Data Transmission

    Test message sending and receiving from each client to verify data integrity.

Error Handling

    Test scenarios like network disconnection or invalid credentials to ensure error handling works as expected.

Additional Notes

    Modularity: Each component is modularized to allow for independent execution and testing.
    Security: The application includes basic authentication to restrict unauthorized access.
    Edge Cases: The program is designed to handle basic edge cases like network interruptions and invalid credentials.

"# mesh-network" 
