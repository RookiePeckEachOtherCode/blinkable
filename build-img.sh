#!/bin/bash
services=("base" "api" "user" "homepage" "article")
dockerfiles=("deployment/server/base/Dockerfile" "deployment/server/api/Dockerfile" "deployment/server/user/Dockerfile" "deployment/server/homepage/Dockerfile" "deployment/server/article/Dockerfile")

for ((i=0; i<${#services[@]}; i++)); do
    service=${services[$i]}
    dockerfile=${dockerfiles[$i]}
    
    echo "Building ${service} image..."
    
    if ! docker build -t "blink-${service}:v1" --file "$(pwd)/${dockerfile}" .; then
        echo "Failed to build ${service} image"
        exit 1
    fi
done

echo "All images built successfully"
