FROM golang:1.17
WORKDIR "/go/src/github.com/f4tal-err0r/discordbot"
COPY *.go ./
COPY go.* ./
COPY config/* ./config/
RUN go mod download
RUN go build -o /app/discordbot .
WORKDIR /app
CMD [ "./discordbot" ]