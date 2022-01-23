FROM golang:1.17-alpine
WORKDIR /app
COPY *.go ./
COPY go.* ./
RUN go mod download
RUN go build -o discordbot
CMD [ "./discordbot" ]