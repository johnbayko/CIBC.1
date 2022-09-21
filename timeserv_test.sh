#!/bin/zsh

port=0
if [ "$#" -ge 1 ]
then
    port=$1
else
    read "port?Port:"
fi

echo port ${port}

echo GET time
curl http://localhost:${port}/api/time

echo GET time tz=abc
curl "http://localhost:${port}/api/time?tz=abc"

echo GET time tz=abc, def
curl "http://localhost:${port}/api/time?tz=abc,def"
