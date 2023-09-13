# Latest golang image on apline linux
FROM golang:1.21.0-alpine3.18

# Env variables
ENV GOOS linux
ENV CGO_ENABLED 0
ENV GOPROXY https://proxy.golang.com.cn,direct

# Work directory
WORKDIR /magellanic-l

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Building the application
RUN go build -o magellanic-l ./cmd/api/

# Starting our application
# CMD ["go", "run", "main.go"]
CMD ./magellanic-l

# Exposing server port
EXPOSE 9999
