FROM golang

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/PlantWaterMe/GardenMonitor

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN go build -o queue-consumer ./cmd/queue-consumer/main.go

# Run the executable
CMD ["./queue-consumer"]