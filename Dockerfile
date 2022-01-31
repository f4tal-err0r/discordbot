FROM golang:1.17
WORKDIR /app
COPY *.go ./
COPY go.* ./
COPY config/ ./config/
COPY hiscore/ ./hiscore/
COPY .git/ ./.git/
RUN go mod download
RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
    go build -ldflags "-X main.GitCommit=$GIT_COMMIT" -o /app/discordbot .

CMD [ "./discordbot" ]