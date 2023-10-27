FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./

#copy main.go && http.go for compile it inside container
COPY ./cmd/user/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./user-api
RUN chmod +x ./user-api
RUN chmod +x ./wait-for-it.sh
EXPOSE 8088

#will be overrided from docker-compose environment
ENV MYSQL_HOST=localhost
ENV MYSQL_PORT=3306
ENV WAITFORIT_TIMEOUT=30

CMD ./wait-for-it.sh $MYSQL_HOST:$MYSQL_PORT -- ./user-api

#CMD ./user-api