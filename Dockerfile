# Latest golang image on apline linux
FROM 1.21.0-alpine3.18

# Env variables
ENV GOOS linux
ENV CGO_ENABLED 0

# Work directory
WORKDIR /magellanic-l

# Installing dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copying all the files
COPY . .

# Building the application
RUN go build -o magellanic-l

# Starting our application
# CMD ["go", "run", "main.go"]
CMD ./docker-go

# Exposing server port
EXPOSE 9999
