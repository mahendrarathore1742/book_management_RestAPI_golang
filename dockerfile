FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /Book_management_system

EXPOSE 3000

CMD [ "/Book_management_system" ]