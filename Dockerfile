FROM golang:1.15
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o /app/bin


FROM scratch
COPY --from=0 /app/bin .
CMD ["./bin"]