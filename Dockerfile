# # Get base image of golang
# FROM golang:1.18.1

# # # Define WORKDIR path
# RUN mkdir /app
# ADD . /app
# WORKDIR /app

# # # Download dependancies
# RUN go mod download && go mod verify

# # # Build the file to the given output
# RUN go build -o main .

# # # Run the executable
# CMD [ "/app/main" ]

#----------------------------------------------------------------

# Get base image of golang
FROM golang:1.18.1-alpine3.15

RUN apk --no-cache add gcc g++ make git

# # Define WORKDIR path
RUN mkdir /app
ADD . /app
WORKDIR /app

# # Download dependancies
RUN go mod download && go mod verify

# # Build the file to the given output
RUN go build -o main .

# # Run the executable
CMD [ "/app/main" ]

#----------------------------------------------------------------


#  Multi stage - NOT WORKING PROPERLY
# FROM golang:1.18.1-alpine AS build
# RUN apk --no-cache add gcc g++ make git
# ADD . /app
# WORKDIR /app
# RUN go mod download && go mod verify
# RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go

# FROM alpine:3.13
# RUN apk --no-cache add ca-certificates
# WORKDIR /usr/bin
# COPY --from=build /app/bin /go/bin
# EXPOSE 5555
# ENTRYPOINT /go/bin/web-app --port 5555


#----------------------------------------------------------------
# FROM golang:1.18.1
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download && go mod verify
# COPY . .
# RUN go install github.com/cespare/reflex@latest
# EXPOSE 5555
# CMD reflex -g '*.go' go run main.go -s

#----------------------------------------------------------------
