# Start from a Debian image with the latest version of Go installed
# (GOPATH) configured at /go.
FROM golang:latest
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o rainbow .
EXPOSE 3000
CMD ["/app/rainbow"]
