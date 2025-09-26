# Net-Cat

## Overview

Welcome to **TCPChat**! This project is a simple recreation of the classic **NetCat** utility in a **Server-Client** architecture using **Go**. The goal is to create a group chat where multiple clients can connect to a server, send messages, and receive messages from each other in real time. It supports TCP connections and offers essential features like message history, timestamping, and user identification.

---

## Features

- **Multiple Clients**: The server supports multiple clients, allowing a 1-to-many relationship (one server, many clients).
- **Client Identification**: Clients must provide a non-empty name when connecting to the server.
- **Timestamped Messages**: All messages are prefixed with a timestamp and the sender's name.
- **Message History**: New clients joining the chat will receive all previous messages.
- **Join/Leave Notifications**: The server notifies all clients when someone joins or leaves the chat.
- **No Empty Messages**: Clients are not allowed to send empty messages.
- **Port Specification**: The default port is `8989`, but clients can specify a custom port if needed.
- **Connection Limit**: The server can handle a maximum of 10 clients at any given time.

---

## Team

This project was developed by
 - **Jawad Zahraoui**
 - **Yassine Bourazza**.

---

## Requirements

- **Programming Language**: Go
- **Concurrency**: The project uses **Go-routines** to handle concurrent client connections.
- **Synchronization**: Channels or **Mutexes** are used for synchronization.
- **Error Handling**: Proper error handling on both the server and client sides.
- **Testing**: Unit tests for both the server and client to ensure everything works as expected.

---

## Allowed Packages

- `io`
- `log`
- `os`
- `fmt`
- `net`
- `sync`
- `time`
- `bufio`
- `errors`
- `strings`
- `reflect`

---

## Usage

### Starting the Server

To start the server and listen on the default port `8989`, run the following:

```bash
$ ./TCPchat 2525
Server Listening on port: 2525
