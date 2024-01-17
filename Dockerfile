FROM golang:1.21

WORKDIR /src

COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -o /bin/app .

CMD ["/bin/app"]
