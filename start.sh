#!/bin/bash

# build 
 docker-compose build

# start MySQL
docker-compose up mysql -d

# init database
./init-mysql.sh

# start demo service
docker-compose up app -d
