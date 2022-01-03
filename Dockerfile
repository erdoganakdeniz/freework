FROM golang:1.17.2


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /freework

EXPOSE 8080

CMD [ "/freework" ]