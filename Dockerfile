FROM golang:alpine AS build-env

ADD $PWD/app /app

# Create a "nobody" non-root user for the next image by crafting an /etc/passwd
# file that the next image can copy in. This is necessary since the next image
# is based on scratch, which doesn't have adduser, cat, echo, or even sh.
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

# Build the Go app with CGO_ENABLED=0 so we use the pure-Go implementations for
# things like DNS resolution (so we don't build a binary that depends on system
# libraries)
RUN cd /app && CGO_ENABLED=0 go build -o goapp

# final stage
FROM scratch

COPY --from=build-env /app/goapp /app/
COPY --from=build-env /etc_passwd /etc/passwd

USER nobody
ENTRYPOINT ["/app/goapp"]
