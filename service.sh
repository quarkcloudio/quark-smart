#!/bin/bash

if [ "$1"x == "start"x ]; then
    bash "./scripts/start_service.sh"
fi

if [ "$1"x == "stop"x ]; then
    bash "./scripts/stop_service.sh"
fi