#
# Dockerfile
#
# docker build --pull -t gaelgirodon/propencrypt .
# docker tag gaelgirodon/propencrypt gaelgirodon/propencrypt:$VERSION
# docker push gaelgirodon/propencrypt:$VERSION
# docker push gaelgirodon/propencrypt:latest

FROM golang:1.22 AS build
COPY . /app
WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o ./propencrypt ./cmd/propencrypt.go

FROM buildpack-deps:stable-curl
COPY --from=build /app/propencrypt /usr/local/bin/
CMD ["propencrypt"]
