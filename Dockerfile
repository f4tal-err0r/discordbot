FROM golang:1.16-alpine
WORKDIR /app
COPY *.go ./
COPY go.* ./
RUN go build -o discordbot
CMD [ "./discordbot" ]