# net-cat

A simple chat server that allows multiple users to connect and chat with each other.

## Usage

To start the server, navigate to the project directory and run the following command:

`go run . [port number]`

Allowed port numbers are digits from 0 - 65535

If no port number is provided, the server will default to port 8989.

To connect to the server, you can use a telnet client and connect to the server's IP address and port number.

## Commands

- `/name`: Change your name
- `/exit`: Exit the server
- `/help`: Show the available commands

## Notes

- The server can only handle a maximum of 10 concurrent connections.

## Dependencies

- Go 1.15 or newer

## Kudos for inspiration

- Victor
- Jay

## Authors <a name = "authors"></a>

- [@Falusvampen](https://github.com/Falusvampen)
