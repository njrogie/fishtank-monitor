#!/bin/bash

#remove all previous traces of setup
docker stop $(docker ps -q)
docker rm $(docker ps -aq)
docker volume rm roach
docker image rm roach
docker network rm mynet

docker volume create roach

docker network create -d bridge mynet

docker run -d --name roach --hostname db --network mynet -p 26257:26257 -p 8080:8080 \
    -v  roach:/cockroach/cockroach-data cockroachdb/cockroach:latest-v20.1 start-single-node \
    --insecure

docker exec roach ./cockroach sql --insecure --execute "CREATE DATABASE mydb;"
docker exec roach ./cockroach sql --insecure --execute "CREATE USER nick;"
docker exec roach ./cockroach sql --insecure --execute "GRANT ALL ON DATABASE mydb TO nick"
docker exec roach ./cockroach sql --insecure --execute "CREATE TABLE IF NOT EXISTS mydb.ftdata (\
    EntryTime TIMESTAMP PRIMARY KEY,\
    Tank_ID int,\
    Temperature FLOAT,\
    Ph FLOAT)"

