FROM golang:1.18 as build
WORKDIR /app
COPY . .
RUN APP_NAME=api make build 
RUN APP_NAME=cli make build

FROM alpine:3.16
EXPOSE 80
CMD ["/opt/api"]
COPY --from=build /app/bin/api /opt/api
COPY --from=build /app/bin/cli /opt/cli

