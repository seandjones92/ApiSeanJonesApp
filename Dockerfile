#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go install -v ./...

# #final stage
FROM scratch
COPY --from=builder /go/bin/apiseanjonesapp /apiseanjonesapp
CMD [ "/apiseanjonesapp" ]
EXPOSE 8080