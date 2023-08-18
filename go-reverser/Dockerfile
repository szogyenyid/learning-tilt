FROM golang:1.21
WORKDIR /app
ADD ./src/ .
RUN go install ./
ENTRYPOINT php-with-tilt