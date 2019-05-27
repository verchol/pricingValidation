FROM golang:1.12.5-alpine3.9
COPY . /src
WORKDIR /src
RUN go get 
RUN go build -o /bin/main
CMD ["/bin/main"]