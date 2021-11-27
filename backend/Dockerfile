FROM golang:1.16-alpine
ENV PORT 8080

WORKDIR /app/backend
# RUN go install

ENV DB_TYPE postgres
ENV DB_HOST db
ENV DB_PORT 5432
ENV DB_USER root
ENV DB_PASSWORD root
ENV DB_NAME commusto
ENV DB_SCHEMA commusto
ENV DB_SSL disable

# install go hot reroad tool
RUN go get -u github.com/cosmtrek/air
EXPOSE 8080

# CMD ["air"]