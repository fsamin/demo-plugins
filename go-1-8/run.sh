#!/bin/bash

echo '######################'
echo '#### Building app ####'
echo '######################'
docker build -t go18-plugin .

echo '######################'
echo '#### Running app  ####'
echo '######################'
docker run -it --rm --name go18-plugin go18-plugin
