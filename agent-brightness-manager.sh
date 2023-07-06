#!/bin/bash

# Define the path to the Go program
go_program="./swayidle-brightness-manager"

# Define a timer in seconds
timer=30

# Launch swayidle with desired commands
swayidle \
    timeout $timer "$go_program dim" \
    resume "$go_program restore" \
    before-sleep 'swaylock' &

wait
