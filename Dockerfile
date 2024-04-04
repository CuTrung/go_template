
## Build Stage ##
FROM golang:1.22.1-alpine as build_stage

ENV APP_NAME go_template

WORKDIR $APP_NAME

COPY . .

RUN CGO_ENABLED=0 go build -v -o /$APP_NAME/bin/ .

## Run Stage ##
FROM alpine:3.14

ENV APP_NAME go_template

# COPY .env .
COPY --from=build_stage /$APP_NAME/bin/ .

EXPOSE 3000

CMD ./$APP_NAME