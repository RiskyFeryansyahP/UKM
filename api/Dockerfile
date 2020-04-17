# Stage 1 - build the application

# FROM defines the base image used to start the build process.
FROM golang:1.14-alpine as builder
# RUN is the central executing directive for Dockerfiles.
RUN mkdir /build
# ADD copies the files from a source on the host into the container’s own filesystem at the set destination.
ADD . /build/
# WORKDIR sets the path where the command, defined with CMD, is to be executed.
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Stage 2 - copy build executable from builder to alpine
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
# EXPOSE associates a specific port to enable networking between the container and the outside world.
EXPOSE 5000
# CMD can be used for executing a specific command within the container.
CMD ["./main", "--migrate=true"]