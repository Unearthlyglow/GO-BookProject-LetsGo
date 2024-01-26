# syntax=docker/dockerfile:1


##
## Build the application from source
##


FROM golang:1.21 AS build-stage


WORKDIR /app/cmd  


#Adjusted WORKDIR to the subdirectory


COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .  
# Copy the entire project (including cmd directory)

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-muse-registry ./cmd  
# Specify the cmd directory

##
## Run the tests in the container
##

FROM build-stage AS run-test-stage
RUN go test -v ./...

##
## Deploy the application binary into a lean image
##

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-muse-registry /sveltego  
# Updated binary path

EXPOSE 8080


USER nonroot:nonroot

ENTRYPOINT ["/sveltego"]  
# Adjusted binary name

