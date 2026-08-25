#using Linux env with go installed
FROM golang:1.21

#working dir inside the container
WORKDIR /app

#copying my go file from the machine into container
COPY main.go .

#compile the code
RUN go build -o my-program main.go

#execute my program
CMD ["./my-program"]