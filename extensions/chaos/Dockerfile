FROM golang:1.15 as build-image

WORKDIR /go/src

COPY go.mod ./
RUN go mod download

COPY main.go ./
COPY extension extension
COPY proxy proxy
RUN CGO_ENABLED=0 go build -o ../bin

FROM public.ecr.aws/lambda/go:1

COPY --from=build-image /go/bin/ /opt/extensions/
