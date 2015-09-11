#!/usr/bin/env bash
echo "start program at : `date` "
nohup ./mblog > /dev/null 2>&1 &
echo $! > pid