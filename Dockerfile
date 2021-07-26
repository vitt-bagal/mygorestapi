 
FROM golang:alpine AS build
RUN apk update && \
    apk add curl \
            git \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*

# download and build it 
RUN git clone https://github.com/vitt-bagal/mygorestapi.git \
    && cd mygorestapi \
    && go mod download && go mod verify \
    && go build -o /bin/server .
# Base image 
FROM alpine:latest
# Default supliers URL
ENV FRUIT_SUPPLIER="https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b"
ENV VEG_SUPPLIER="https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c"
ENV GRAIN_SUPPLIER="https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148"
RUN apk --no-cache add ca-certificates bash

# Copy built in binary 
COPY --from=build /bin/server /bin/server

# Default Port used for server
EXPOSE 9091

CMD ["/bin/server"]
# End of Dockerfile