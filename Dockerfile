FROM golang
WORKDIR /icebaby-user-service
COPY . /icebaby-user-service
RUN go build main.go 
EXPOSE 9090
ENTRYPOINT ./main -m release