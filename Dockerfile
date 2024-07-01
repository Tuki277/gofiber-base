FROM golang:1.21-alpine AS builder

LABEL maintainer="huyth"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

FROM alpine

EXPOSE 8000

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/apiserver", "/build/.env", "/"]
RUN mkdir database
RUN mkdir static
COPY --from=builder ["/build/database/GeoLite2-Country.mmdb", "/database"]
COPY --from=builder ["/build/database/GeoLite2-City.mmdb", "/database"]
COPY --from=builder ["/build/static/", "/static/"]

RUN mkdir logs


# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]