#!/bin/sh

if [ "$1" == "start"x ]; then
    sh "./scripts/start_service.sh"
fi

if [ "$1" == "stop"x ]; then
    sh "./scripts/stop_service.sh"
fi