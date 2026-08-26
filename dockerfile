#using Linux env with go installed
FROM golang:1.21

#working dir inside the container
WORKDIR /app

#copying my go file from the machine into container
COPY network.go .

#compile the code
RUN go build -o server network.go

#execute my program
CMD ["./server"]