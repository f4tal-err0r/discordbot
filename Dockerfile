FROM golang:1.16-alpine
WORKDIR /app
COPY * ./
RUN go build -o discordbot
CMD [ "./discordbot" ]