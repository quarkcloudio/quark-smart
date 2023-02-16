#!/bin/sh

if [ "$1"x == "start"x ]; then
    sh "./scripts/start_service.sh"
fi

if [ "$1"x == "stop"x ]; then
    sh "./scripts/stop_service.sh"
fi