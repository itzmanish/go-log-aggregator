#!/bin/bash

count=1
echo "infinite loops [ hit CTRL+C to stop]" 
for (( ; ; ))
do
   echo "log entry $count"
   echo "log entry $count" >> sample/log.txt
   ((count+=1))
   sleep 0.1;
done