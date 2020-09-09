FROM golang:1.15 as build

WORKDIR /go/src/fiber-boilerplate
COPY . .

RUN go mod download
RUN go mod verify
RUN GOOS=linux go build -o app

FROM golang:1.15 as run

ENV APP_HOME /go/src/fiber-boilerplate

WORKDIR $APP_HOME

COPY --from=build $APP_HOME/app $APP_HOME

EXPOSE 3000
CMD ./app serve