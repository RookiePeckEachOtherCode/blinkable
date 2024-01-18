#!/bin/bash
docker-compose -f ./docker-compose.yml up -d
docker-compose -f ./deployment/nacos/nacos-docker/standalone-mysql-8.yaml up -d
