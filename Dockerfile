FROM golang:1.15

# Set the working directory to the root of the project
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o net-cat .

# Set the port number that the server will listen on
ENV PORT 8989

# Run the server when the container starts
CMD ["./net-cat"]
