#!/bin/bash

IMAGE_NAME=farport/http-redirect-go:0.1
IMAGE_ID=$(docker images -qf"reference=$IMAGE_NAME")

if ! [ -z "$IMAGE_ID" ]; then
    docker rmi $IMAGE_ID
fi

docker build -t $IMAGE_NAME .
