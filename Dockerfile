FROM golang:1.17-alpine AS build
WORKDIR /app/
COPY . .
RUN go install
# RUN CGO_ENABLED=0 go test ./...
RUN go build -o /bin/pool example/*.go
EXPOSE 3000

FROM alpine:3.13.2
COPY --from=build /bin/pool /bin/pool
ENV PORT=80
EXPOSE 80
ENTRYPOINT ["/bin/pool"]