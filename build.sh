#!/usr/bin/env bash


IMAGE_NAME=ufp/env-file

echo "PART 1 BUILD IT"
docker build -t ${IMAGE_NAME} .

echo "PART 2 TEST IT"



echo "PART 3 RUN IT"

docker run --rm ${IMAGE_NAME}
