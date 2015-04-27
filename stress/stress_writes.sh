#!/bin/bash

THREADS=$1
CONNECTIONS=$2
DURATION=$3
URL=$4

docker run -it --rm \
  --net=host \
  -v `pwd`:/data williamyeh/wrk:4.0.1 \
  --threads=$THREADS \
  --connections=$CONNECTIONS \
  --duration=$DURATION \
  --script=request_counter.lua \
  $URL
