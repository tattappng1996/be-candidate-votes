FROM golang:1.19-buster as builder

WORKDIR /be-candidate-votes
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o ./be-candidate-votes .

FROM scratch
ARG HTTP_PORT
ENV HTTP_PORT=:8081

COPY --from=builder ./be-candidate-votes ./

ENTRYPOINT [ "./be-candidate-votes" ]