FROM golang:1.15 as build-image

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 go build -o ../bin

FROM extension-chaos as chaos-extension
FROM public.ecr.aws/lambda/go:1

COPY --from=chaos-extension /opt/extensions/ /opt/extensions/
COPY --from=build-image /go/bin/ /var/task/

# Command can be overwritten by providing a different command in the template directly.
CMD ["hello-world"]
