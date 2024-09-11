FROM golang:1.9-alpine 

WORKDIR /app

COPY . .

RUN go mod tidy 
RUN go build -o main ./cmd/api 

EXPOSE 8080

CMD [ "./main" ]