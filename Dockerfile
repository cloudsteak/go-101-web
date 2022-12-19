FROM golang:1.19.4-alpine


# Set the Current Working Directory inside the container
WORKDIR /


# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -u .

# Install the package
RUN go install -v .


# Build package
RUN go build -o /go-101-web

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["go-101-web"]