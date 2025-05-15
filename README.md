# Go gRPC Project

This project demonstrates the implementation of gRPC services in Go, featuring multiple service types and MongoDB integration. It includes examples of various gRPC communication patterns and secure connections using TLS.

## Services

### 1. Greet Service
- Unary RPC: Basic greeting service
- Server Streaming: `GreetManyTimes`
- Client Streaming: `LongGreet`
- Bidirectional Streaming: `GreetEveryone`
- Deadline/Timeout: `GreetWithDeadline`

### 2. Calculator Service
- Unary RPC: Sum operation
- Server Streaming: Prime number decomposition
- Client Streaming: Compute average
- Bidirectional Streaming: Find maximum
- Error Handling: Square root with validation

### 3. Blog Service (with MongoDB)
- Create Blog
- Read Blog
- Update Blog
- Delete Blog
- List Blogs

## Prerequisites

1. Go 1.16 or higher
2. Protocol Buffers
3. MongoDB (for blog service)
4. Make

## Installation

1. Clone the repository:
```bash
git clone https://github.com/thiendsu2303/go-grpc.git
cd go-grpc
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate protobuf code:
```bash
make greet
make calculator
make blog
```

## SSL/TLS Configuration

The project includes TLS support. To generate certificates:

1. Navigate to the SSL directory:
```bash
cd ssl
```

2. Run the certificate generation script:
```bash
./ssl.sh
```

## Running the Services

### Greet Service
```bash
# Start server
./bin/greet/server

# Run client
./bin/greet/client
```

### Calculator Service
```bash
# Start server
./bin/calculator/server

# Run client
./bin/calculator/client
```

### Blog Service
1. Start MongoDB (using Docker):
```bash
cd blog
docker-compose up -d
```

2. Run the service:
```bash
# Start server
./bin/blog/server

# Run client
./bin/blog/client
```

## Project Structure

```
.
├── calculator/       # Calculator service implementation
├── greet/           # Greeting service implementation
├── blog/            # Blog service with MongoDB
├── ssl/             # SSL/TLS certificates
└── Makefile         # Build and generate commands
```

## Features

- Multiple gRPC patterns (Unary, Server Streaming, Client Streaming, Bidirectional Streaming)
- SSL/TLS security
- Error handling
- Deadlines and timeouts
- MongoDB integration
- Docker support for database
- Makefile for easy building

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.