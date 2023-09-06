FROM golang

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/PlantWaterMe/GardenMonitor

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

RUN go build -o waterapi ./cmd/api/main.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./waterapi"]