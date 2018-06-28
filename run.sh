#!/bin/sh
i=1
while : ; do
  
  ./app
  echo "deleted $i"
  sleep 1
  i=$(($i+1))
done
echo "deleted all $i"
