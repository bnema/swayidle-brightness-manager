# Swayidle Brightness Manager

This tiny Go application is a utility for Hyprland or Sway that get your actual brightness and set it to 10% after X seconds of idle time. The bash script also locks the screen before the system goes to sleep. It is designed to work with `swayidle` and `swaylock`.


## Features

- Decreases the screen brightness to 10% after 30 seconds of idle time.
- Restores the brightness to its original level upon resuming activity.
- Ensures that the screen is locked before the system goes to sleep.
- Prevent dimming the screen if the user is watching a video.

## Pre-requisites

- `go 1.20` or higher
- `Sway or Hyprland`
- `swayidle` 
- `swaylock` 
- `brightnessctl`
- `playerctl`

## Service example

```bash

[Unit]
Description=Sway Idle Brightness Manager

[Service]
Environment=PATH=/home/YOURUSERNAME/.local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ExecStart=/bin/bash %h/.local/bin/agent-brightness-manager.sh
Restart=always

[Install]
WantedBy=default.target

```