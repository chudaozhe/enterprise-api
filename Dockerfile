FROM golang:1.19 AS build
WORKDIR /usr/src/app
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -mod=mod --tags netgo -v -o /usr/local/bin/app .
# CMD ["app"]

FROM alpine:3.16
WORKDIR "/data"
EXPOSE 7097
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=build /usr/local/bin/app .
CMD ["./app"]