# In the first stage, we want to import files and build a golang binary.
#FROM golang:1.19.3-buster AS build
 FROM golang:buster AS build

 WORKDIR /app

 COPY go.mod ./
 COPY go.sum ./

 RUN go mod download && go mod verify

 COPY main.go ./

 RUN go build -o /my-app

#In the second stage, we take a distroless image and copy our binary there.
#For the local development, you can just use a single stage and perhaps find a way to reload and rebuild the image automatically on any change.
# for Debug See: https://github.com/GoogleContainerTools/distroless#debug-images
 FROM gcr.io/distroless/base-debian11

 COPY --from=build /my-app /my-app

 ENTRYPOINT ["/my-app"]
