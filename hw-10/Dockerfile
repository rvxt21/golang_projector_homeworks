FROM golang

WORKDIR /api

COPY . .

RUN go build -o api .

CMD [ "./api" ]