FROM golang:1.22.5-alpine

# Install necessary packages
RUN apk add --no-cache ca-certificates curl make

# Install Atlas for database migrations
RUN curl -sSfL https://atlasgo.sh | sh

# Set the Current Working Directory inside the container
WORKDIR /app

COPY ["Makefile", "./"]

# Copy the pre-built binary into the container
COPY ["bin/projmgmt", "projmgmt"]

COPY migrate migrate

EXPOSE 80

CMD ["./projmgmt"]
