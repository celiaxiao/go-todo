#!/usr/bin/env bash

# serverPID=`ps -ef | grep "go run" | grep -v grep | awk '{print $2}' `
bindingPID=`lsof -i :8080 |awk '{print $2}' `

for serverPID in `ps -ef|grep 'go run' |grep -v "grep"|awk '{print $2}'`
do
    if [ $serverPID != "PID" ]; then
        kill $serverPID
    fi
	# kill -9 $proclist
done

for serverPID in `lsof -i :8080 |awk '{print $2}'`
do
    if [ $serverPID != "PID" ]; then
        kill $serverPID
    fi
	# kill -9 $proclist
done
# echo $serverPID $bindingPID
# kill $serverPID $bindingPID