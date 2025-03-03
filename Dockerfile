FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ="Asia/Jakarta"

ARG COMMIT=""

WORKDIR /root/

# Copy the Pre-built binary
# the binary should be run by
COPY platform-usergroup .

# Not Secret env variable
ENV DEBUG_MODE=true
ENV REST_SERVER="0.0.0.0:80"
ENV SERVICE_SERVER="0.0.0.0:40"
ENV APP_SERVICE="liat-platform-usergroup:40"
ENV APP_VERSION=${COMMIT}

# Expose port 80 to the outside world
EXPOSE 80 40

# Command to run the executable
CMD ["./platform-usergroup"]