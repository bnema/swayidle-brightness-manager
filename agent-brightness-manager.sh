#!/bin/bash

# Define the path to the Go program
go_program="./SwayIdleBrightLock"

# Launch swayidle with desired commands
swayidle \
    timeout 10 "$go_program dim" \
    resume "$go_program restore" \
    before-sleep 'swaylock' &

wait