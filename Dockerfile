FROM golang:alpine AS build

# cache dependencies (disabled until there are non-stdlib dependencies)
#ADD go.mod go.sum ./
#RUN go mod download

# build
WORKDIR /app
ADD main.go go.mod /app/
RUN go build -o /main

# copy artifacts to a clean image
FROM golang:alpine
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]
