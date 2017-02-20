#!/bin/bash

PI_IP="10.0.0.160"

if [ -z $1 ]; then
  echo "No file provided."
  echo "Example: ./piload.sh example.txt"
  exit 1
fi

scp ./$1 pi@$PI_IP:~/
exit
