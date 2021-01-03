# First stage: build the executable.
FROM golang:buster AS builder

# It is important that these ARG's are defined after the FROM statement
ARG SSH_PRIV="nothing"
ARG SSH_PUB="nothing"
ARG GOSUMDB=off

# Create the user and group files that will be used in the running
# container to run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'finger:x:65534:65534:finger:/:' > /user/passwd && \
    echo 'finger:x:65534:' > /user/group

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/shitpostingio/fingerprint-microservice

# Import the code from the context.
COPY .  .

# Install libraries
RUN apt-get update; \
    apt install -y libavcodec-dev libavformat-dev libavutil-dev libswscale-dev xz-utils

# Build the executable
RUN go install

# Finale stage
FROM golang:buster

# Install libraries
RUN apt update; \
    apt install -y libavcodec-dev libavformat-dev libavutil-dev libswscale-dev xz-utils

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Copy the built executable
COPY --from=builder /go/bin/fingerprint-microservice /home/finger/fingerprint

RUN chown -R finger /home/finger

# Set the workdir
WORKDIR /home/finger

# Expose port 10000
EXPOSE 10000

# Perform any further action as an unprivileged user.
USER finger:finger

# Run the compiled binary.
CMD ["./fingerprint"]
