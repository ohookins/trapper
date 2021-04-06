FROM golang:alpine 

# cache dependencies
ADD go.mod go.sum ./
RUN go mod download

# build
ADD . .
RUN go build -o /main

# copy artifacts to a clean image
FROM golang:alpine
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]
