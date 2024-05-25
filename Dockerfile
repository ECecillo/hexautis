FROM golang:1.22.3 as build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -v -o bin/hexautis ./cmd/main.go

# Use a Docker multi-stage build to create a lean production image.

FROM alpine:latest

CMD ["/bin/sh"]

RUN adduser -D -u 12345 -g 12345 hexautis

# Copy the binary to the production image from the builder stage.
COPY --from=build /app/bin/hexautis /usr/bin/hexautis

USER hexautis

WORKDIR /home/hexautis

# Run
ENTRYPOINT ["hexautis"]
