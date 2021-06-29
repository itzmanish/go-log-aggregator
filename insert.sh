#!/bin/bash

count=0
echo "infinite loops [ hit CTRL+C to stop]" 
for (( ; ; ))
do
   echo "log entry $count" >> sample/log.txt
   ((count+=1))
   sleep 1;
done