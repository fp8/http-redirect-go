#!/bin/bash

IMAGE_NAME=farport/http-redirect-go:0.1

docker run --rm -p8080:8080 -it $IMAGE_NAME
