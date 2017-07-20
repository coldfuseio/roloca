#!/bin/bash

docker build -t registry.gitlab.com/coldfuse-oss/roloca .
docker push registry.gitlab.com/coldfuse-oss/roloca
