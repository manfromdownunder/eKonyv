FROM node:16 as web
WORKDIR /web
COPY . .
WORKDIR /web/server/app/
RUN npm install
RUN npm run build

FROM golang as build
WORKDIR /go/src/
COPY . .
COPY --from=web /web/ .

ENV CGO_ENABLED=0
RUN go get -d -v ./...
RUN go install -v ./...
WORKDIR /go/src/cmd/eKonyv/
RUN go build

FROM gcr.io/distroless/static as app
WORKDIR /app
COPY --from=build /go/src/cmd/eKonyv/eKonyv .

EXPOSE 80
VOLUME [ "/books" ]
ENV BASE_PATH=/

ENTRYPOINT ["./eKonyv", "server", "--dir", "/books", "--port", "80"]
