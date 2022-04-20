# Get base image of golang
FROM golang

# Define WORKDIR path
RUN mkdir /app
ADD . /app
WORKDIR /app

# Download dependancies
RUN go mod download && go mod verify

# Build the file to the given output
RUN go build -o main .

# Run the executable
CMD [ "/app/main" ]