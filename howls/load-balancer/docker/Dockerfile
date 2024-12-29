FROM golang:1.23.1 AS base

WORKDIR /app
#COPY ../go.mod ../go.sum  ./
#RUN go mod download
COPY ../. .
RUN go mod download


FROM base AS worker
RUN go build -o worker-server ./worker
EXPOSE 8080
ENTRYPOINT ["sh", "-c", "export WORKER_ID=$(hostname | awk -F'-' '{print $NF}') && exec \"$@\"", "--"]
CMD ["./worker-server"]


FROM base AS load-balancer
RUN go build -o load-balancer-server ./load-balancer
EXPOSE 8080
CMD ["./load-balancer-server"]
