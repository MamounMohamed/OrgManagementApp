FROM golang:alpine
WORKDIR /App
ENV GOPROXY=https://proxy.golang.org,direct

COPY . .
RUN go mod download 
RUN go build -o main .

CMD ["./main"]
EXPOSE 8080



