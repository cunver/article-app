#FROM golang:1.12.0-alpine3.9
FROM golang:1.17
RUN mkdir /article-app
ADD . /article-app
WORKDIR /article-app
RUN go build -o main .
CMD ["/article-app/main"]
