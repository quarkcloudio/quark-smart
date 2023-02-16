#!/bin/sh

start=start
stop=stop

if [ $0 = $start ]; then
    sh "./scripts/start_service.sh"
fi

if [ $0 = $stop ]; then
    sh "./scripts/stop_service.sh"
fi