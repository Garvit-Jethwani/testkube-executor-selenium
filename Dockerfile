# syntax=docker/dockerfile:1
FROM golang:1.18
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0 
ENV GOOS=linux
RUN cd cmd/agent;go build -o /runner -mod mod -a .

FROM node:18-alpine3.14
RUN apk add --no-cache --update \
    npm \
    git   
RUN npm install -g mocha
WORKDIR /home/runner
COPY --from=0 /runner /bin/runner
# ENV DEPENDENCY_MANAGER=npm
ENTRYPOINT ["/bin/runner"]