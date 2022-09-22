#!/bin/zsh

port=0
if [ "$#" -ge 1 ]
then
    port=$1
    echo port ${port}
else
    read "port?Port:"
fi

echo GET time
curl http://localhost:${port}/api/time

echo GET time tz=America/New_York
curl "http://localhost:${port}/api/time?tz=America/New_York"

echo GET time tz=America/New_York, Asia/Kolkata
curl "http://localhost:${port}/api/time?tz=America/New_York,Asia/Kolkata"

echo GET time tz=X
curl "http://localhost:${port}/api/time?tz=X"

echo GET time tz=X, Asia/Kolkata
curl "http://localhost:${port}/api/time?tz=X,Asia/Kolkata"

echo GET time tz=America/New_York, X
curl "http://localhost:${port}/api/time?tz=America/New_York,X"

