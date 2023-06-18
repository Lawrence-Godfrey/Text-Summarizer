#!/bin/bash

IMAGE_NAME="text-summarizer"
REPO_URI=<repo_uri>


docker build --platform linux/amd64 -t $IMAGE_NAME .

docker tag $IMAGE_NAME:latest $REPO_URI:latest

aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin $REPO_URI

docker push $REPO_URI:latest
