#!/bin/sh

if [ $0 -eq "start" ]; then
    sh "./scripts/start_service.sh"
fi

if [ $0 -eq "stop" ]; then
    sh "./scripts/stop_service.sh"
fi